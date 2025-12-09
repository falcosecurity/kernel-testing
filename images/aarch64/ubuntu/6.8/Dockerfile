FROM ubuntu:24.04

ARG VERSION=6.8.0-91
ARG URL='http://ports.ubuntu.com/ubuntu-ports/pool/main/l/linux/'

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
		grep -v "64k" | \
		cut -d\" -f8 | \
		xargs -I@ curl -LO ${URL}@ && \
	curl ${URL} | \
		grep -E "linux-headers-${VERSION}|linux-tools-common-${VERSION}" | \
		grep -E "_all.deb" | \
		grep -v "64k" | \
		cut -d\" -f8 | \
		xargs -I@ curl -LO ${URL}@ && \
	find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
	rm -f *.deb && \
	sed -ie '/^ConditionVirtualization.*/d' /lib/systemd/system/systemd-timesyncd.service && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id  && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
