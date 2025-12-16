FROM amazonlinux:2023

ARG VERSION=6.1.34-58
ARG URL='https://cdn.amazonlinux.com/al2023/core/mirrors/latest/x86_64'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN dnf groupinstall -y 'Development Tools' && \
	dnf install -y \
		binutils-devel \
		clang \
		iproute \
		kmod \
		libcap-devel \
		llvm \
		make \
		openssh-server \
		rsync \
		systemd && \
	curl -sLO "$URL"/mirror.list && \
	URL=$(cat mirror.list) && \
	curl -sLO "${URL}"repodata/primary.xml.gz && \
	gunzip primary.xml.gz && \
    grep -E 'kernel-headers-'${VERSION}'.*.rpm' primary.xml | grep href | cut -d\"  -f2 | \
		xargs -I@ curl -Lo headers.rpm ${URL}@ && \
    grep -E 'kernel-devel-'${VERSION}'.*.rpm' primary.xml | grep href | cut -d\"  -f2 | \
		xargs -I@ curl -Lo sources.rpm ${URL}@ && \
	rpm -ivh --force --nodeps ./headers.rpm ./sources.rpm && \
	rm -vf ./headers.rpm ./sources.rpm && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	mkdir -p /lib/modules/${VERSION}.102.amzn2023.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.102.amzn2023.x86_64/ /lib/modules/${VERSION}.102.amzn2023.x86_64/build && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
