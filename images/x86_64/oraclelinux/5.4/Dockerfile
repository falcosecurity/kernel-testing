FROM oraclelinux:8

ARG VERSION=5.4.17-2136.320.7.1
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL8/UEKR6/x86_64/'

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
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	curl -LO https://cmake.org/files/v3.16/cmake-3.16.7-Linux-x86_64.tar.gz && \
	tar -xaf cmake-3.16.7-Linux-x86_64.tar.gz && \
	rm -vf cmake-3.16.7-Linux-x86_64.tar.gz && \
	mv cmake-3.16.7-Linux-x86_64   /opt/cmake-3.16.7 && \
	ln -sf /opt/cmake-3.16.7/bin/* /usr/bin/ && \
	mkdir -p /lib/modules/5.4.17-2136.320.7.1.el8uek.x86_64/ && \
	ln -s /usr/src/kernels/5.4.17-2136.320.7.1.el8uek.x86_64/ /lib/modules/5.4.17-2136.320.7.1.el8uek.x86_64/build && \
	yum clean all && \
	rm -rf /var/cache/yum
