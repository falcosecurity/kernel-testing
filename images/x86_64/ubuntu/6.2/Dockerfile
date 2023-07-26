FROM ubuntu:22.04

ARG VERSION=6.2.0-26-generic

WORKDIR /home/ubuntu

RUN echo 'deb http://mirrors.edge.kernel.org/ubuntu/ jammy main' > /etc/apt/sources.list.d/kernel.org.list && \
	apt-get update && apt-get install -y \
		build-essential \
		clang \
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
	echo "root:root" | chpasswd
