FROM oraclelinux:7

ARG VERSION=3.10.0-1160.92.1.0.2
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL7/MODRHCK/x86_64'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

COPY /dev.repo /etc/yum.repos.d/

RUN yum install -y \
		binutils-devel \
		gcc \
		git \
		iproute \
		kmod \
		libcap-devel \
		make \
		openssh-server \
		rsync \
		systemd \
		tar && \
	curl -Lo kernel.rpm  ${URL}/getPackage/kernel-${VERSION}.el7.x86_64.rpm && \
	curl -Lo devel.rpm   ${URL}/getPackage/kernel-devel-${VERSION}.el7.x86_64.rpm && \
	yum install -y ./*.rpm  && \
	rm -f ./*.rpm  && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	mkdir -p /lib/modules/${VERSION}.el7.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el7.x86_64/ /lib/modules/${VERSION}.el7.x86_64/build && \
	yum install -y oracle-softwarecollection-release-el7 && \
	yum install -y llvm-toolset-7 && \
	find /opt/rh/llvm-toolset-7/root/usr/bin/ ! -type d -exec ln -s {} /usr/bin \; && \
	yum clean all && \
	rm -rf /var/cache/yum
