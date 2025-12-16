FROM ubuntu:24.04

ARG VERSION=6.8.0-91
ARG URL='http://mirrors.edge.kernel.org/ubuntu/pool/main/l/linux/'

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
	rm -rf /var/lib/apt/lists/* && \
	curl ${URL} | \
		grep -E "linux-headers-${VERSION}|linux-tools-${VERSION}|linux-tools-common_${VERSION}" | \
		cut -d\" -f2 | \
		xargs -I@ curl -LO ${URL}@ && \
	find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
	rm -f *.deb && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
