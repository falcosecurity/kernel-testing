FROM fedora:33

ARG VERSION=5.8.15-301
ARG HVERSION=5.8.11-300
ARG URL='https://archives.fedoraproject.org/pub/archive/fedora/linux/releases/33/Everything/x86_64/os/Packages'

RUN dnf groupinstall -y 'Development Tools' && \
	dnf install -y \
		clang \
		cmake \
		iproute \
		libcap-devel \
		llvm \
		openssh-server \
		rsync \
		systemd && \
	curl -Lo sources.rpm ${URL}/k/kernel-devel-${VERSION}.fc33.x86_64.rpm    && \
	curl -Lo headers.rpm ${URL}/k/kernel-headers-${HVERSION}.fc33.x86_64.rpm && \
	dnf install -y ./*.rpm && \
	rm -v ./*.rpm && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	mkdir -p /lib/modules/${VERSION}.fc33.x86_64 && \
	ln -s /usr/src/kernels/${VERSION}.fc33.x86_64/ /lib/modules/${VERSION}.fc33.x86_64/build && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
