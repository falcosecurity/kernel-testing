FROM oraclelinux:8

ARG VERSION=5.4.17-2136.320.7.1
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL8/UEKR6/x86_64/'
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
	curl -Lo kernel.rpm ${URL}/getPackage/kernel-uek-${VERSION}.el8uek.x86_64.rpm && \
	curl -Lo devel.rpm  ${URL}/getPackage/kernel-uek-devel-${VERSION}.el8uek.x86_64.rpm && \
	yum install -y ./*.rpm  && \
	rm -vf ./*.rpm  && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	mkdir -p /lib/modules/${VERSION}.el8uek.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el8uek.x86_64/ /lib/modules/${VERSION}.el8uek.x86_64/build && \
	yum clean all && \
	rm -rf /var/cache/yum
