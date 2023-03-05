# Falco drivers tests

This repository automatically runs Falco [drivers_test](https://github.com/falcosecurity/libs/tree/master/test/drivers) through Ansible.

## Prerequisites

* Install [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
* Install [Ignite](https://ignite.readthedocs.io/en/stable/installation/)

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
ansible-playbook modern-bpf-test.yml --ask-become 
```

## Clean-up all machines

```bash
ansible-playbook clean-up.yml --ask-become
```
