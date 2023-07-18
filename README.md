[![Drivers Matrix Tests](https://github.com/alacuku/e2e-falco-tests/actions/workflows/kernel_tests.yaml/badge.svg)](https://github.com/alacuku/e2e-falco-tests/actions/workflows/kernel_tests.yaml)

# Falco drivers tests

This repository automatically runs Falco [scap-open](https://github.com/falcosecurity/libs/tree/master/userspace/libscap/examples/01-open) binary on all supported drivers through Ansible.

## Prerequisites

* Install [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
* Install [Ignite](https://ignite.readthedocs.io/en/stable/installation/) from therealbobo fork (use main branch): https://github.com/therealbobo/ignite

## Configure repository

Configure an ssh key pair into the `/group_vars/all/vars.yml` file 👇

```yml
#####################
# SSH configuration #
#####################

#Path to the generated SSH private key file
ssh_key_path: "" # <-- Replace here with the key path
ssh_key_name: "" # <-- Replace here with the key name

#Path to the private key
prv_key_path: "{{ssh_key_path}}/{{ssh_key_name}}"

#path to the public key used to ssh to the machines, if this key does not exist then a new one is generated with the same name
pub_key_path: "{{ssh_key_path}}/{{ssh_key_name}}.pub"
```

You need to provide the path to the key pair (`ssh_key_path`) and the name of the key pair (`ssh_key_name`)

## Run tests

From the repository root you can run tests on all machines by typing:

```bash
ansible-playbook master-playbook.yml --ask-become 
```

To rerun tests:

```bash
ansible-playbook scap-open-test.yml --ask-become 
```

## Clean-up all machines

```bash
ansible-playbook clean-up.yml --ask-become
```
