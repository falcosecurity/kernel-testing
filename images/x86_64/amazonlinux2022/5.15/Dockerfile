FROM amazonlinux:2022

ARG VERSION=5.15.73-45
ARG URL='https://al2022-repos-us-east-1-9761ab97.s3.dualstack.us-east-1.amazonaws.com/core/mirrors/latest/x86_64'


RUN dnf groupinstall -y 'Development Tools' && \
	dnf install -y \
		binutils-devel \
		clang \
		cmake \
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
	dnf install -y ./*.rpm && \
	rm -vf ./*.rpm && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	mkdir -p /lib/modules/${VERSION}.135.amzn2022.x86_64/ && \
	ln -s /usr/src/kernels/${VERSION}.135.amzn2022.x86_64/ /lib/modules/${VERSION}.135.amzn2022.x86_64/build && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
