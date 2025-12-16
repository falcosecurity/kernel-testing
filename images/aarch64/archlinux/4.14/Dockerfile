FROM lopsided/archlinux:latest

ARG VERSION=4.14.15-1
ARG URL='http://tardis.tiny-vps.com/aarm/packages'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

WORKDIR /home/ubuntu

RUN pacman -Syyu --noconfirm && \
	pacman -S --noconfirm \
		clang \
		gcc \
		git \
		libcap \
		llvm \
		make \
		openssh \
		python \
		rsync \
		systemd \
		wget && \
	CMAKE_ARCH="aarch64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	ln -sf /usr/share/zoneinfo/US/Eastern /etc/localtime && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.gen && \
	locale-gen && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.conf && \
	echo 'archlinux' > /etc/hostname && \
	systemctl enable sshd && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	curl -Lo headers.tar.xz ${URL}/l/linux-aarch64-headers/linux-aarch64-headers-${VERSION}-aarch64.pkg.tar.xz && \
	pacman -U --noconfirm ./headers.tar.xz && \
	rm -v ./headers.tar.xz && \
	ln -s /usr/lib/systemd/systemd /sbin/init && \
	pacman -Scc --noconfirm || exit 0
