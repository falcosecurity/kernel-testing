FROM ubuntu:18.04

ARG VERSION=4.15.0-1118-aws
ARG URL='https://archive.ubuntu.com/ubuntu/pool/main/l/linux-aws/'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

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
		cut -d\" -f8 | \
		xargs -I@ curl -LO ${URL}@ && \
	find . -name '*.deb' | xargs -n1 dpkg -i --force-depends && \
	rm -f *.deb && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	echo "" > /etc/machine-id && echo "" > /var/lib/dbus/machine-id && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd \
