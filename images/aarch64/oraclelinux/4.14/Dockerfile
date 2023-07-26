FROM oraclelinux:7

ARG VERSION=4.14.35-2047.527.2
ARG URL='https://yum.oracle.com/repo/OracleLinux/OL7/latest/aarch64'

COPY /dev.repo /etc/yum.repos.d/

RUN yum groupinstall -y 'Development Tools' && \
	yum install -y \
                binutils-devel \
                elfutils-libelf-devel \
                git \
                iproute \
                kmod \
                libcap-devel \
                make \
                openssh-server \
                oracle-softwarecollection-release-el7 \
                rsync \
                systemd \
                tar && \
	yum install -y oracle-armtoolset-1 && \
	curl -L -o kernel.rpm  ${URL}/getPackage/kernel-uek-${VERSION}.el7uek.aarch64.rpm && \
	curl -L -o devel.rpm   ${URL}/getPackage/kernel-uek-devel-${VERSION}.el7uek.aarch64.rpm && \
	curl -Lo libdtrace-devel.rpm ${URL}/getPackage/libdtrace-ctf-devel-1.1.0-2.el7.aarch64.rpm && \
	curl -Lo libdtrace.rpm ${URL}/getPackage/libdtrace-ctf-1.1.0-2.el7.aarch64.rpm && \
	yum install -y --skip-broken ./kernel.rpm ./devel.rpm ./libdtrace.rpm ./libdtrace-devel.rpm && \
	echo 'source /opt/oracle/oracle-armtoolset-1/enable' >> /root/.bashrc && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	curl -LO https://cmake.org/files/v3.19/cmake-3.19.8-Linux-aarch64.tar.gz && \
	tar -xaf cmake-3.19.8-Linux-aarch64.tar.gz && \
	rm -vf cmake-3.19.8-Linux-aarch64.tar.gz && \
	mv cmake-3.19.8-Linux-aarch64   /opt/cmake-3.19.8 && \
	ln -sf /opt/cmake-3.19.8/bin/* /usr/bin/ && \
	mkdir -p /lib/modules/${VERSION}.el7uek.aarch64/ && \
	ln -s /usr/src/kernels/${VERSION}.el7uek.aarch64/ /lib/modules/${VERSION}.el7uek.aarch64/build && \
	yum clean all && \
	rm -rf /var/cache/yum
