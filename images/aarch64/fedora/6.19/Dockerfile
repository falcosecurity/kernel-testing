FROM fedora:44

ARG VERSION=6.19.7-300
ARG HDR_VERSION=6.19.6-300
ARG ARCHITECTURE=aarch64
# Ensure we use the correct variable name consistently
ARG URL='https://archives.fedoraproject.org/pub/fedora/linux/development/44/Everything/aarch64/os/Packages'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf install --disablerepo=updates-testing -y @development-tools \
        clang \
        iproute \
        libcap-devel \
        llvm \
        openssh-server \
        rsync \
        systemd && \
    curl -Lo devel.rpm "${URL}/k/kernel-devel-${VERSION}.fc44.${ARCHITECTURE}.rpm" && \
    curl -Lo headers.rpm "${URL}/k/kernel-headers-${HDR_VERSION}.fc44.${ARCHITECTURE}.rpm" && \
    dnf install -y ./*.rpm && \
    rm -v ./*.rpm && \
    dnf clean all && \
	rm -rf /var/cache/yum
RUN	mkdir -p /lib/modules/${VERSION}.fc44.${ARCHITECTURE} && \
	ln -s /usr/src/kernels/${VERSION}.fc44.${ARCHITECTURE}/ /lib/modules/${VERSION}.fc44.${ARCHITECTURE}/build
RUN	CMAKE_ARCH="aarch64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf -
RUN	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
