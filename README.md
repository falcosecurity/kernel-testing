[![Falco kernel tests Repository](https://github.com/falcosecurity/evolution/blob/main/repos/badges/falco-infra-blue.svg)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#infra-scope) 
[![Incubating](https://img.shields.io/badge/status-incubating-orange?style=for-the-badge)](https://github.com/falcosecurity/evolution/blob/main/REPOSITORIES.md#incubating)
![Architectures](https://img.shields.io/badge/ARCHS-x86__64%7Caarch64-blueviolet?style=for-the-badge)
[![Latest release](https://img.shields.io/github/v/release/falcosecurity/kernel-testing?style=for-the-badge)](https://github.com/falcosecurity/kernel-testing/releases/latest)

# Falco drivers tests

This repository automatically runs Falco [scap-open](https://github.com/falcosecurity/libs/tree/master/userspace/libscap/examples/01-open) binary on all supported drivers through Ansible, spawning Firecracker microVMs to test Falco drivers against multiple kernels.  
You can find list of machines being used [here](./ansible-playbooks/group_vars/all/vars.yml#L18).

## Prerequisites

* Install [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
* Install [Ignite](https://ignite.readthedocs.io/en/stable/installation/) from `therealbobo` fork (use `main` branch): https://github.com/therealbobo/ignite; just issue `make` and then `sudo make install` to install everything needed under `/usr/local/`.
* Install ignite CNI plugins by following this guide: https://ignite.readthedocs.io/en/stable/installation/#cni-plugins:
```bash
export CNI_VERSION=v0.9.1
export ARCH=$([ $(uname -m) = "x86_64" ] && echo amd64 || echo arm64)
sudo mkdir -p /opt/cni/bin
curl -sSL https://github.com/containernetworking/plugins/releases/download/${CNI_VERSION}/cni-plugins-linux-${ARCH}-${CNI_VERSION}.tgz | sudo tar -xz -C /opt/cni/bin
```

## Configure

It is advised to avoid directly modifying [`vars.yml`](ansible-playbooks/group_vars/all/vars.yml) file;  
instead one can create a local vars.yml file to override keys from the default vars.  

The only mandatory thing to be configured is an ssh key pair:
```yml
#Path to the generated SSH private key file
ssh_key_path: "" # <-- Replace here with the key path
ssh_key_name: "" # <-- Replace here with the key name
```
## Run

From the `ansible-playbooks` directory you can run tests on all machines by typing:

```bash
ansible-playbook main-playbook.yml --ask-become --extra-vars "@/path/to/local/vars.yaml"
```

To rerun tests:

```bash
ansible-playbook scap-open.yml --ask-become --extra-vars "@/path/to/local/vars.yaml"
```

To cleanup all machines

```bash
ansible-playbook clean-up.yml --ask-become --extra-vars "@/path/to/local/vars.yaml"
```

## CI Usage

To better suit the CI usage, a [Github composite action](https://docs.github.com/en/actions/creating-actions/creating-a-composite-action) has been developed.  
Therefore, running kernel-testing in your Github workflow is as easy as adding this step:
```
- uses: falcosecurity/kernel-testing@main
  # Give it an id to be able to later use its outputs
  id: kernel_tests
  with:
    # libs version to be tested, eg: master.
    # Default: 'master'
    libsversion: master
    
    # libs repo to be tested, eg: falcosecurity/libs.
    # Default: 'falcosecurity/libs'
    libsrepo: falcosecurity/libs
    
    # Whether to generate matrixes as matrix artifact.
    # Default: false
    build_matrix: 'true'
```
Then you can use action outputs to retrieve artifacts:
```
- uses: actions/upload-artifact@latest
  with:
    name: ansible_output
    path: ${{ steps.kernel_tests.outputs.ansible_output }}
        
- uses: actions/upload-artifact@latest
  with:
    name: matrix
    path: ${{ steps.kernel_tests.outputs.matrix_output }}
```

As an example, see [libs reusable workflow](https://github.com/falcosecurity/libs/blob/master/.github/workflows/reusable_kernel_tests.yaml).

> __NOTE:__ Since we don't use annotated tags, one cannot use eg: falcosecurity/kernel-testing@v0, but only either exact tag name or master.

> __NOTE:__ Of course, you'll need to run your tests on virtualization-enabled nodes.
