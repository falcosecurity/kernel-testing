FROM oraclelinux:9

ARG VERSION=5.15.0-8.91.4.1
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL9/baseos/latest/aarch64'
ARG ARCHITECTURE='aarch64'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN yum install -y \
		binutils-devel \
		clang \
		gcc \
		git \
		iproute \
		kmod \
		libcap-devel \
		llvm \
		make \
		openssh-server \
		rsync \
		systemd \
		tar && \
	curl -L -o kernel.rpm  ${URL}/getPackage/kernel-uek-core-${VERSION}.el9uek.${ARCHITECTURE}.rpm && \
	curl -L -o devel.rpm   ${URL}/getPackage/kernel-uek-devel-${VERSION}.el9uek.${ARCHITECTURE}.rpm && \
	curl -L -o modules.rpm ${URL}/getPackage/kernel-uek-modules-${VERSION}.el9uek.${ARCHITECTURE}.rpm && \
	yum install -y ./kernel.rpm ./devel.rpm ./modules.rpm && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	CMAKE_ARCH="aarch64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	mkdir -p /lib/modules/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/ && \
	ln -s /usr/src/kernels/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/ /lib/modules/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/build
