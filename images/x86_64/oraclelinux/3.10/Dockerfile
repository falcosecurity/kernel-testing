FROM oraclelinux:7

ARG VERSION=3.10.0-1160.92.1.0.2
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL7/MODRHCK/x86_64'

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
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	curl -LO https://cmake.org/files/v3.16/cmake-3.16.7-Linux-x86_64.tar.gz && \
	tar -xaf cmake-3.16.7-Linux-x86_64.tar.gz && \
	rm -vf cmake-3.16.7-Linux-x86_64.tar.gz && \
	mv cmake-3.16.7-Linux-x86_64   /opt/cmake-3.16.7 && \
	ln -sf /opt/cmake-3.16.7/bin/* /usr/bin/ && \
	mkdir -p /lib/modules/${VERSION}.el7.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el7.x86_64/ /lib/modules/${VERSION}.el7.x86_64/build && \
	yum install -y oracle-softwarecollection-release-el7 && \
	yum install -y llvm-toolset-7 && \
	find /opt/rh/llvm-toolset-7/root/usr/bin/ ! -type d -exec ln -s {} /usr/bin \; && \
	yum clean all && \
	rm -rf /var/cache/yum
