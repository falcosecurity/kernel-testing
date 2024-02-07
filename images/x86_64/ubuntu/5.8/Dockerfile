FROM ubuntu:20.04

ARG VERSION=5.8.0-1041-aws

WORKDIR /home/ubuntu

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
		build-essential \
		clang-7 \
		cmake \
		curl \
		dbus \
		git \
		iproute2 \
		iputils-ping \
		kmod \
		libcap-dev \
		libelf-dev \
		linux-headers-${VERSION} \
		linux-tools-${VERSION} \
		llvm \
		net-tools \
		openssh-server \
		rng-tools \
		rsync \
		sudo \
		systemd \
		udev \
		wget && \
	apt-get clean && \
	rm -rf /var/lib/apt/lists/* && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	ln -s /usr/bin/clang-7 /usr/bin/clang
