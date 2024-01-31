FROM fedora:38

ARG VERSION=6.2.9-300
ARG HVERSION=6.2.6-300
ARG ARCHITECTURE=aarch64
ARG URL='https://mirrors.kernel.org/fedora/releases/38/Everything/aarch64/os/Packages'

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
	curl -Lo sources.rpm ${URL}/k/kernel-devel-${VERSION}.fc38.${ARCHITECTURE}.rpm    && \
	curl -Lo headers.rpm ${URL}/k/kernel-headers-${HVERSION}.fc38.${ARCHITECTURE}.rpm && \
	dnf install -y ./*.rpm && \
	rm -v ./*.rpm && \
	mkdir -p /lib/modules/${VERSION}.fc38.${ARCHITECTURE} && \
	ln -s /usr/src/kernels/${VERSION}.fc38.${ARCHITECTURE}/ /lib/modules/${VERSION}.fc38.${ARCHITECTURE}/build && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
