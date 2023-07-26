FROM ubuntu:18.04

ARG VERSION=4.15.0-213
ARG URL='http://mirrors.edge.kernel.org/ubuntu/pool/main/l/linux/'

WORKDIR /home/ubuntu

RUN apt-get update && apt-get install -y \
		build-essential \
		clang \
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
		grep -vE 'i386|lowlatency' | \
		cut -d\" -f2 | \
		xargs -I@ curl -LO ${URL}@ && \
	find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
	rm -f *.deb && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	curl -LO https://cmake.org/files/v3.16/cmake-3.16.7-Linux-x86_64.tar.gz && \
	tar -xaf cmake-3.16.7-Linux-x86_64.tar.gz && \
	rm -vf cmake-3.16.7-Linux-x86_64.tar.gz && \
	mv cmake-3.16.7-Linux-x86_64   /opt/cmake-3.16.7 && \
	ln -sf /opt/cmake-3.16.7/bin/* /usr/bin/
