FROM centos:centos7.9.2009

ARG VERSION=3.10.0-1160
ARG BASEOS_URL='http://mirror.centos.org/centos/7.9.2009/os/x86_64/Packages'

RUN yum group install -y "Development Tools" && \
	yum install -y \
		cmake \
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
	mkdir -p /lib/modules/${VERSION}.el7.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el7.x86_64/ /lib/modules/${VERSION}.el7.x86_64/build && \
	curl -LO https://cmake.org/files/v3.16/cmake-3.16.7-Linux-x86_64.tar.gz && \
	tar -xaf cmake-3.16.7-Linux-x86_64.tar.gz && \
	rm -vf cmake-3.16.7-Linux-x86_64.tar.gz && \
	mv cmake-3.16.7-Linux-x86_64   /opt/cmake-3.16.7 && \
	ln -sf /opt/cmake-3.16.7/bin/* /usr/bin/ && \
	yum clean all && \
	rm -rf /var/cache/yum && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
