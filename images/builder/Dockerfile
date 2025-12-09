FROM centos:7

# TARGETARCH is automatically set by BuildKit (arm64 or amd64)
ARG TARGETARCH

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
		cmake \
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
        sed && \
    yum -y install devtoolset-9-gcc devtoolset-9-gcc-c++ && \
	source /opt/rh/devtoolset-9/enable

RUN	curl -L -o /tmp/cmake.tar.gz https://github.com/Kitware/CMake/releases/download/v3.22.5/cmake-3.22.5-linux-$(uname -m).tar.gz && \
	gzip -d /tmp/cmake.tar.gz && \
	tar -xpf /tmp/cmake.tar --directory=/tmp && \
	cp -R /tmp/cmake-3.22.5-linux-$(uname -m)/* /usr && \
	rm -rf /tmp/cmake-3.22.5-linux-$(uname -m) cmake.tar && \
	yum clean all && \
	rm -rf /var/cache/yum /usr/man /usr/doc && \
	systemctl enable dhclient && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
