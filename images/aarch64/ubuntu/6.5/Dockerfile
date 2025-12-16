FROM ubuntu:22.04

ARG VERSION=6.5.0-1024-aws
ARG URL='https://ports.ubuntu.com/ubuntu-ports/pool/main/l/linux-aws-6.5/'

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

WORKDIR /home/ubuntu

RUN apt-get update && apt-get install -y \
		build-essential \
		clang \
		cmake \
		curl \
		curl \
		dbus \
		git \
		iproute2 \
		iputils-ping \
		kmod \
		libcap-dev \
		libelf-dev \
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
	curl ${URL} | \
		grep -E "linux-headers-${VERSION}|linux-tools-${VERSION}" | \
		grep "arm64" | \
		cut -d\" -f8 | \
		xargs -I@ curl -LO ${URL}@ && \
	find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
	rm -f *.deb && \
	sed -ie '/^ConditionVirtualization.*/d' /lib/systemd/system/systemd-timesyncd.service && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id  && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
