FROM busybox:stable AS initrd-base

FROM ubuntu:22.04

WORKDIR /home/ubuntu/

RUN mkdir -p /home/ubuntu/out/

COPY --from=initrd-base / /home/ubuntu/out/

ADD create.sh Makefile.virtio_mmio /opt/initrd-builder/

RUN apt update && \
	apt install -y \
		build-essential \
		libelf-dev \
		gcc \
		kmod \
		xz-utils \
		p7zip-full \
		rpm2cpio \
		curl \
		cpio \
		binutils \
		bzip2 \
		gzip \
		lzop \
		lz4 \
		zstd && \
	apt-get clean && \
	rm -rf /var/lib/apt/lists/* && \
	curl -Lo /usr/bin/extract-vmlinux https://raw.githubusercontent.com/torvalds/linux/master/scripts/extract-vmlinux && \
	chmod +x /usr/bin/extract-vmlinux
