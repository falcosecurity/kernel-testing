FROM centos:7

# TARGETARCH is automatically set by BuildKit (arm64 or amd64)
ARG TARGETARCH
ARG CMAKE_VERSION

# Use bash for all RUN steps with some safety options enabled.
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

COPY /dhclient.service /usr/lib/systemd/system/

# Fix broken mirrors - centos:7 eol
RUN sed -i s/mirror.centos.org/vault.centos.org/g /etc/yum.repos.d/*.repo; \
    sed -i s/^#.*baseurl=http/baseurl=https/g /etc/yum.repos.d/*.repo; \
    sed -i s/^mirrorlist=http/#mirrorlist=https/g /etc/yum.repos.d/*.repo

RUN yum -y install centos-release-scl

# fix broken mirrors (again)
RUN sed -i s/mirror.centos.org/vault.centos.org/g /etc/yum.repos.d/*.repo; \
    sed -i s/^#.*baseurl=http/baseurl=https/g /etc/yum.repos.d/*.repo; \
    sed -i s/^mirrorlist=http/#mirrorlist=https/g /etc/yum.repos.d/*.repo && \
	if [ "$TARGETARCH" != "amd64" ]; then \
        sed -i 's|vault.centos.org/centos/7/sclo|vault.centos.org/altarch/7/sclo|g' /etc/yum.repos.d/CentOS-SCLo-scl*.repo; \
    fi

RUN yum install -y \
		curl \
		dhclient \
		gcc \
		git \
		iproute \
		libcap-devel \
		libstdc++ \
		m4 \
		make \
		openssh-server \
		rpm-build \
		rsync \
		systemd \
		tar \
	    wget \
        sed \
        devtoolset-9-gcc \
        devtoolset-9-gcc-c++

RUN	CMAKE_RELEASE_NAME="cmake-${CMAKE_VERSION}-linux-$(uname -m)" && \
	curl -L "https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/${CMAKE_RELEASE_NAME}.tar.gz" | \
		tar -C /usr --strip-components 1 "${CMAKE_RELEASE_NAME}/bin" "${CMAKE_RELEASE_NAME}/share" -xzpvf - && \
	yum clean all && \
	rm -rf /var/cache/yum /usr/man /usr/doc && \
	systemctl enable dhclient && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
