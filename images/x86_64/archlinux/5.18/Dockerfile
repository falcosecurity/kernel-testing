FROM archlinux/archlinux:base-20220807.0.72894

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

WORKDIR /home/ubuntu

RUN printf 'SigLevel = Never\nServer=https://archive.archlinux.org/repos/2022/08/04/$repo/os/$arch' > /etc/pacman.d/mirrorlist && \
pacman -Syyu --noconfirm && \
	pacman -S --noconfirm \
		clang \
		cmake \
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
	ln -sf /usr/share/zoneinfo/US/Eastern /etc/localtime && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.gen && \
	locale-gen && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.conf && \
	echo 'archlinux' > /etc/hostname && \
	systemctl enable sshd && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
