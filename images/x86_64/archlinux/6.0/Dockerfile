FROM archlinux:base

ARG VERSION=6.0.12.arch1-1
ARG URL='https://archive.archlinux.org/packages'

WORKDIR /home/ubuntu

RUN pacman -Syyu --noconfirm && \
	pacman -S --noconfirm \
		clang \
		cmake \
		gcc \
		git \
		libcap \
		llvm \
		make \
		openssh \
		python \
		rsync \
		wget && \
	yes | pacman -Scc && \
	curl -Lo headers.tar.zst ${URL}/l/linux-headers/linux-headers-${VERSION}-x86_64.pkg.tar.zst && \
	curl -Lo bpf.tar.zst     ${URL}/b/bpf/bpf-6.0-2-x86_64.pkg.tar.zst && \
	pacman -U --noconfirm ./headers.tar.zst && \
	pacman -U --noconfirm ./bpf.tar.zst && \
	rm -v ./headers.tar.zst ./bpf.tar.zst && \
	ln -sf /usr/share/zoneinfo/US/Eastern /etc/localtime && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.gen && \
	locale-gen && \
	echo 'LANG=en_US.UTF-8' > /etc/locale.conf && \
	echo 'archlinux' > /etc/hostname && \
	systemctl enable sshd && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
