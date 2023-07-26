FROM amazonlinux:2023

ARG VERSION=6.1.34-58
ARG URL='https://cdn.amazonlinux.com/al2023/core/mirrors/latest/x86_64'


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
    grep -E 'bpftool-'${VERSION}'.*.rpm' primary.xml | grep href | cut -d\"  -f2 | \
		xargs -I@ curl -Lo bpftool.rpm ${URL}@ && \
	dnf install -y ./headers.rpm ./bpftool.rpm ./sources.rpm && \
	rm -vf ./headers.rpm ./bpftool.rpm ./sources.rpm && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	mkdir -p /lib/modules/6.1.34-58.102.amzn2023.x86_64/ && \
	ln -s /usr/src/kernels/6.1.34-58.102.amzn2023.x86_64/ /lib/modules/6.1.34-58.102.amzn2023.x86_64/build && \
	sed -i -e 's/^AcceptEnv LANG LC_\*$/#AcceptEnv LANG LC_*/' /etc/ssh/sshd_config && \
	echo 'UseDNS no' >> /etc/ssh/sshd_config && \
	echo "root:root" | chpasswd
