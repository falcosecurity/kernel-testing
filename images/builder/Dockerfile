FROM centos:7

COPY /dhclient.service /usr/lib/systemd/system/

RUN yum -y install centos-release-scl && \
	yum -y install devtoolset-9-gcc devtoolset-9-gcc-c++ && \
	source /opt/rh/devtoolset-9/enable && \
	yum install -y \
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
		wget && \
	curl -L -o /tmp/cmake.tar.gz https://github.com/Kitware/CMake/releases/download/v3.22.5/cmake-3.22.5-linux-$(uname -m).tar.gz && \
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
