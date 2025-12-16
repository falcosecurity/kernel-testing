FROM archlinux/archlinux:base-20220807.0.72894

ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

WORKDIR /home/ubuntu

RUN printf 'SigLevel = Never\nServer=https://archive.archlinux.org/repos/2022/08/04/$repo/os/$arch' > /etc/pacman.d/mirrorlist && \
pacman -Syyu --noconfirm && \
	pacman -S --noconfirm \
		clang \
		gcc \
		git \
		libcap \
		linux-headers \
		llvm \
		make \
		openssh \
		python \
		rsync \
		wget && \
	pacman -Scc --noconfirm && \
	CMAKE_ARCH="x86_64" && \
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
	echo "root:root" | chpasswd
