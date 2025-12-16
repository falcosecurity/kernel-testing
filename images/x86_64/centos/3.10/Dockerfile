FROM centos:centos7.9.2009

ARG VERSION=3.10.0-1160
ARG BASEOS_URL='https://vault.centos.org/centos/7.9.2009/os/x86_64/Packages'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# Fix broken mirrors.
RUN sed -i s/mirror.centos.org/vault.centos.org/g /etc/yum.repos.d/*.repo; \
    sed -i s/^#.*baseurl=http/baseurl=https/g /etc/yum.repos.d/*.repo; \
    sed -i s/^mirrorlist=http/#mirrorlist=https/g /etc/yum.repos.d/*.repo

RUN yum group install -y "Development Tools" && \
	yum install -y \
		curl \
		gcc \
		git \
		iproute \
		libcap-devel \
		libstdc++ \
		make \
		openssh-server \
		rsync \
		systemd \
		tar && \
	curl -Lo headers.rpm ${BASEOS_URL}/kernel-devel-${VERSION}.el7.x86_64.rpm && \
	yum install -y ./headers.rpm && \
	rm -f ./headers.rpm && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	mkdir -p /lib/modules/${VERSION}.el7.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el7.x86_64/ /lib/modules/${VERSION}.el7.x86_64/build && \
	yum clean all && \
	rm -rf /var/cache/yum && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
