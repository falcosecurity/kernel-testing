FROM amazonlinux:2

ARG VERSION=5.4.247-162.350
ARG URL='http://amazonlinux.us-east-1.amazonaws.com/2/extras/kernel-5.4/latest/x86_64'

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
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd && \
	curl -LO https://cmake.org/files/v3.16/cmake-3.16.7-Linux-x86_64.tar.gz && \
	tar -xaf cmake-3.16.7-Linux-x86_64.tar.gz && \
	rm -vf cmake-3.16.7-Linux-x86_64.tar.gz && \
	mv cmake-3.16.7-Linux-x86_64   /opt/cmake-3.16.7 && \
	ln -sf /opt/cmake-3.16.7/bin/* /usr/bin/ && \
	mkdir -p  /lib/modules/${VERSION}.amzn2.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.amzn2.x86_64/ /lib/modules/${VERSION}.amzn2.x86_64/build && \
	rm -vf /usr/bin/gcc /usr/bin/g++ && \
	update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc10-gcc 100 --slave /usr/bin/g++ g++ /usr/bin/gcc10-c++ && \
	update-alternatives --auto gcc && \
	yum clean all && \
	rm -rf /var/cache/yum
