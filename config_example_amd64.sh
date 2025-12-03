#!/bin/bash

# NOTICE: This script is not intended to be run as is: it is just demonstrative, and is required for the user to go
# through it and adapt the different parts to the specific environment.

### Install miscellaneous dependencies

sudo apt update -y
sudo apt install -y git iproute2 dnsmasq e2tools e2fsprogs ca-certificates curl nano iputils-ping rsync

### Install ansible ###
sudo apt update -y
sudo apt install -y software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install -y ansible-core=2.16.3-0ubuntu2

### Install ansible requirements globally ###
git clone https://github.com/kernel-testing/kernel-testing.git
sudo mkdir -p /usr/share/ansible/collections
sudo chmod 755 /usr/share/ansible/collections
sudo ansible-galaxy collection install -r kernel-testing/requirements.yml -p /usr/share/ansible/collections
echo 'ANSIBLE_COLLECTIONS_PATHS=/usr/share/ansible/collections:$ANSIBLE_COLLECTIONS_PATHS' | sudo tee -a /etc/environment

ansible-galaxy install -r kernel-testing/requirements.yml
sudo ansible-galaxy install -r kernel-testing/requirements.yml

### Install firecracker ###
curl -LO https://github.com/firecracker-microvm/firecracker/releases/download/v1.13.1/firecracker-v1.13.1-x86_64.tgz
tar -xzf firecracker-v1.13.1-x86_64.tgz
sudo mv release-v1.13.1-x86_64/firecracker-v1.13.1-x86_64 /usr/local/bin/firecracker
sudo chmod +x /usr/local/bin/firecracker

### Install docker ###
sudo apt install -y ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update -y
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin

### Install go globally ###
curl -LO https://go.dev/dl/go1.25.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.25.4.linux-amd64.tar.gz
echo 'PATH="/usr/local/go/bin:$PATH"' | sudo tee /etc/environment

### Configure networking ###

## Enable IP forwarding
sudo sysctl -w net.ipv4.ip_forward=1
echo "net.ipv4.ip_forward = 1" | sudo tee /etc/sysctl.d/99-firecracker.conf

## Disable reverse path filtering
CONFIG_FILE="/etc/sysctl.d/99-rp_filter.conf"
sudo bash -c "cat > $CONFIG_FILE" <<'EOF'
net.ipv4.conf.all.rp_filter = 0
net.ipv4.conf.default.rp_filter = 0
EOF
sudo sysctl --system
# just to be super sure that reverse path filtering is disabled for the current interfaces
for f in /proc/sys/net/ipv4/conf/*/rp_filter; do #
    echo 0 | sudo tee "$f"
done

## Configure iptables
# note: pay attention to use the right interface (in place of eth0)
sudo iptables -t nat -A POSTROUTING -s 172.16.0.0/16 -o eth0 -j MASQUERADE
sudo iptables -I FORWARD 1 -s 172.16.0.0/16 -j ACCEPT
sudo iptables -I FORWARD 2 -d 172.16.0.0/16 -j ACCEPT
sudo iptables -I INPUT 1 -i tap+ -j ACCEPT
