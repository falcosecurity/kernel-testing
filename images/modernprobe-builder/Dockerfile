FROM fedora:latest

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

COPY /dhclient.service /usr/lib/systemd/system/

RUN dnf install -y \
		autoconf \
		automake \
		automake \
		bpftool \
		ca-certificates \
		clang \
		cmake \
		dhclient \
		gcc \
		gcc-c++ \
		git \
		iproute \
		libbpf-devel \
		make \
		openssh-server \
		pkg-config \
		rsync \
		systemd && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	systemctl enable dhclient && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
