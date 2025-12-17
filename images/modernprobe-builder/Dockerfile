FROM fedora:39

ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

COPY /dhclient.service /usr/lib/systemd/system/

RUN dnf install -y \
		autoconf \
		automake \
		automake \
		bpftool \
		ca-certificates \
		clang \
		dhclient \
		gcc \
		gcc-c++ \
		git \
		iproute \
		libbpf-devel \
		make \
		openssh-server \
		pkg-config \
		rsync \
		systemd && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-$(uname -m)" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	systemctl enable dhclient && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
