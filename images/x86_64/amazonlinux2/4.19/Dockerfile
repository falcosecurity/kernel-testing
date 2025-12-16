FROM amazonlinux:2

ARG VERSION=4.19.84-33.70
ARG URL='http://amazonlinux.us-east-1.amazonaws.com/2/extras/kernel-ng/latest/x86_64'
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

RUN yum install -y \
		binutils-devel \
		clang \
		gcc10 \
		gcc10-c++ \
		git \
		iproute \
		kmod \
		libcap-devel \
		llvm \
		make \
		openssh-server \
		rsync \
		systemd \
		tar && \
	curl -sLO "$URL"/mirror.list && \
	URL=$(cat mirror.list) && \
	curl -sLO "${URL}"/repodata/primary.xml.gz && \
	gunzip primary.xml.gz && \
    grep -E 'kernel-headers-'${VERSION}'.*.rpm' primary.xml | grep href | cut -d\"  -f2 | \
		xargs -I@ curl -Lo headers.rpm ${URL}/@ && \
    grep -E 'kernel-devel-'${VERSION}'.*.rpm' primary.xml | grep href | cut -d\"  -f2 | \
		xargs -I@ curl -Lo sources.rpm ${URL}/@ && \
	yum install -y ./*.rpm && \
	rm -v ./*.rpm && \
	CMAKE_ARCH="x86_64" && \
	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-${CMAKE_ARCH}" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	mkdir -p /lib/modules/${VERSION}.amzn2.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.amzn2.x86_64/ /lib/modules/${VERSION}.amzn2.x86_64/build && \
	rm -vf /usr/bin/gcc /usr/bin/g++ && \
	update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc10-gcc 100 --slave /usr/bin/g++ g++ /usr/bin/gcc10-c++ && \
	update-alternatives --auto gcc && \
	yum clean all && \
	rm -rf /var/cache/yum
