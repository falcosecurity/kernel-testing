FROM quay.io/centos/centos:stream8

ARG VERSION=4.18.0-497
ARG BASEOS_URL='http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages'

RUN dnf groupinstall -y 'Development Tools' && \
	dnf install -y --allowerasing \
		clang \
		cmake \
		curl \
		git \
		iproute \
		libcap-devel \
		llvm \
		openssh-server \
		rsync \
		systemd && \
	curl -Lo headers.rpm ${BASEOS_URL}/kernel-devel-${VERSION}.el8.x86_64.rpm && \
	curl -Lo bpftool.rpm ${BASEOS_URL}/bpftool-${VERSION}.el8.x86_64.rpm && \
	dnf install -y ./*.rpm && \
	rm -v ./*.rpm && \
	mkdir -p /lib/modules/${VERSION}.el8.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.el8.x86_64/ /lib/modules/${VERSION}.el8.x86_64/build && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	mkdir -p /lib/modules/${VERSION}.fc38.x86_64 && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
