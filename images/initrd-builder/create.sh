#!/bin/bash

LIBPATH=${1:-.}
OUTDIR=/home/ubuntu/out

MODULES=( virtio_mmio virtio_blk virtio_net ext4 crc32c )
UNAMER=$(ls ${LIBPATH}/lib/modules/ | head -n1)

NEWMOD='../virtio_mmio.ko'
OLDMOD=$(find ${LIBPATH} -name virtio_mmio.ko)

if [[ -f ${NEWMOD} ]]; then
	>&2 echo 'Copying virtio_mmio new module'
	cp -vf ${NEWMOD} ${OLDMOD} || {
		mkdir -p ${OUTDIR}/lib/modules/${UNAMER}/kernel/drivers/virtio/
		cp -vf ${NEWMOD} ${OUTDIR}/lib/modules/${UNAMER}/kernel/drivers/virtio/
	}
fi

>&2 echo LIBPATH=$LIBPATH
>&2 echo UNAMER=$UNAMER
depmod -b ${LIBPATH} ${UNAMER}

mkdir -p ${OUTDIR}/lib/modules/${UNAMER}
#cp -a ${LIBPATH}/lib/modules/${UNAMER}/* ${OUTDIR}/lib/modules/${UNAMER}/

cd $LIBPATH

readarray MODULES_PATH < <(
	for MODULE in ${MODULES[@]}; do
		>&2 echo listing deps of module $MODULE
		modprobe -d . -S ${UNAMER} -n --show-depends ${MODULE} | grep -Ev '^builtin' | cut -d.  -f2-
	done
)
for MODULE_PATH in ${MODULES_PATH[@]}; do
	mkdir -p ${OUTDIR}/$(dirname ${MODULE_PATH})
	echo cp ./$MODULE_PATH ${OUTDIR}/$(dirname ${MODULE_PATH})
	cp ./$MODULE_PATH ${OUTDIR}/$(dirname ${MODULE_PATH})
done

mkdir -p ${OUTDIR}/{bin,dev,etc,lib,lib64,mnt,proc,root,sbin,sys}
cp -a /dev/{null,console,tty} ${OUTDIR}/dev/

cat <<EOF > ${OUTDIR}/init
#!/bin/busybox sh

mount -t proc none /proc
mount -t sysfs none /sys
mount -t devtmpfs dev /dev

# Load essential modules
depmod

if grep -q NOMODULES /proc/cmdline ; then
	echo "skipping kmods"
else

EOF

for MODULE in ${MODULES[@]}; do
	if [[ $MODULE == 'virtio_mmio' ]]; then
		echo -e "\tmodprobe $MODULE"' $(cat /proc/cmdline | grep -o "device=[^\ ]*" | paste -s)'
	else
		echo -e "\tmodprobe $MODULE"
	fi
done >> ${OUTDIR}/init

cat <<EOF >> ${OUTDIR}/init

fi

# Mount the root filesystem.
mount /dev/vda /mnt/

# TODO: add support for static ip from cmdline
# e.g.: ifconfig eth0 10.0.2.15
ifconfig eth0 up
udhcpc -t 5 -q -s /bin/simple.script

if [ -L /mnt/etc/resolv.conf ]; then
	unlink /mnt/etc/resolv.conf
	cat /etc/resolv.conf > /mnt/etc/resolv.conf
fi

if grep -q DEBUG /proc/cmdline ; then
	exec /bin/sh
else
	# Clean up.
	umount /proc
	umount /sys

	# Boot the real thing.
	exec switch_root /mnt/ /sbin/init
fi
EOF

chmod +x ${OUTDIR}/init

curl -L -o ${OUTDIR}/bin/simple.script https://git.busybox.net/busybox/plain/examples/udhcp/simple.script
chmod +x ${OUTDIR}/bin/simple.script

cd ${OUTDIR}

find . -print0 | cpio --null --create --verbose --format=newc > initrd
cd -

