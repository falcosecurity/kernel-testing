FROM oraclelinux:9

ARG VERSION=5.15.0-8.91.4.1
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL9/baseos/latest/aarch64'
ARG ARCHITECTURE='aarch64'

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
	curl -LO https://cmake.org/files/v3.19/cmake-3.19.8-Linux-aarch64.tar.gz && \
	tar -xaf cmake-3.19.8-Linux-aarch64.tar.gz && \
	rm -vf cmake-3.19.8-Linux-aarch64.tar.gz && \
	mv cmake-3.19.8-Linux-aarch64   /opt/cmake-3.19.8 && \
	ln -sf /opt/cmake-3.19.8/bin/* /usr/bin/ && \
	mkdir -p /lib/modules/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/ && \
	ln -s /usr/src/kernels/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/ /lib/modules/5.15.0-8.91.4.1.el9uek.${ARCHITECTURE}/build
