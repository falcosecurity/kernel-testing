FROM ubuntu:22.04

ARG VERSION=6.2.0-1018-aws
ARG URL='https://archive.ubuntu.com/ubuntu/pool/main/l/linux-aws-6.2/'

WORKDIR /home/ubuntu

RUN apt-get update && apt-get install -y \
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
	curl ${URL} | \
		grep -E "linux-headers-${VERSION}|linux-tools-${VERSION}|linux-tools-common_${VERSION}" | \
		cut -d\" -f8 | \
		xargs -I@ curl -LO ${URL}@ && \
		find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
		rm -f *.deb && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
