FROM oraclelinux:9

ARG VERSION=5.15.0-8.91.4.1
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL9/UEKR7/x86_64'
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
	curl -Lo kernel.rpm  ${URL}/getPackage/kernel-uek-core-${VERSION}.el9uek.x86_64.rpm && \
	curl -Lo devel.rpm   ${URL}/getPackage/kernel-uek-devel-${VERSION}.el9uek.x86_64.rpm && \
	curl -Lo modules.rpm ${URL}/getPackage/kernel-uek-modules-${VERSION}.el9uek.x86_64.rpm && \
	yum install -y ./*.rpm && \
	rm -f ./*.rpm && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	mkdir -p /lib/modules/${VERSION}.el9uek.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el9uek.x86_64/ /lib/modules/${VERSION}.el9uek.x86_64/build && \
	yum clean all && \
	rm -rf /var/cache/yum
