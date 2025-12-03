# Architecture

This document describes requirements and implementation details of the solution.

## VM spawning

### Requirements

Each VM requires 3 elements to be spawned:

1) a kernel binary (`vmlinux` file)
2) an initramfs (`initrd` file)
3) a `rootfs` ext4 raw disk image

### Implementation

`vmlinux` and `initrd` are extracted from the same corresponding `*-kernel*` docker image, while the `rootfs` is shipped
in a separate `*-image*` docker image. Extraction happens at runtime, and extracted artifacts are cached and reused for
later runs. On each run, a single `rootfs` ext4 disk image is CoW-cloned (shallow copy), and the ephemeral clone is
patched to enable SSHing from the host.

## Networking

### Requirements

- the host must be able to SSH into VMs
- each VM must be able to connect to Internet, to download needed dependencies

### Implementation

Each VM is connected to the host through a TAP interface. For each VM, a `/30` subnet, taken from the `172.16.0.0/16`
range, is allocated.
The subnet is uniquely identified by the `run_id` and the machine name (as specified in
[vars.yml](./ansible-playbooks/group_vars/all/vars.yml)).
For each VM, the first address of the corresponding subnet is assigned to the TAP interface, while the second one is
assigned to the guest OS interface.
Each VM receives its networking configuration through DHCP. The networking configuration includes:

- the guest interface IP address
- the default route (i.e.: the TAP interface IP address)
- DNS configuration (i.e.: `1.1.1.1`)

A dedicated DHCP server (`dnsmasq` instance) is spawned for each TAP interface, specifically configured to offer the
above configuration.
The `dnsmsq` instances are spawned through `systemd`. Given an interface `tapX`, and the corresponding assigned subnet
`172.16.Y.Z/30`, the corresponding `systemd` service instance exposing the DHCP service will be named as follows:

```
`dnsmasq-tap@tapX:172.16.Y.{ Z + 1 }:172.16.Y.{ Z + 2 }.service`
```

Given the aforementioned requirements, on the host:

- IP forwarding must be enabled for IPv4
- reverse path filtering must be disabled on all interfaces
- traffic sourced from `172.16.0.0/16` and exiting the host external interface must be NATted
- the FORWARD chain must allow both incoming and outgoing traffic for `172.16.0.0/16`
- the INPUT chain must allow incoming traffic on `tap+` interfaces`
