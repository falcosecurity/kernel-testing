DRY_RUN := false
REPOSITORY := falcosecurity/kernel-testing
ARCH ?= $(shell uname -m)
YAML_FILE := images.yaml

.PHONY: build-rootfs build-kernel docker-push generate-yaml build-all initrd-builder modernprobe-builder builder

builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building builder image: docker build -t $(REPOSITORY)/builder:0.0.1-$(ARCH) builder"; \
	else \
		echo "Building modernprobe-builder image"; \
		docker build -t $(REPOSITORY)/builder:0.0.1-$(ARCH) builder; \
	fi

modernprobe-builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building modernprobe-builder image: docker build -t $(REPOSITORY)/modernprobe-builder:0.0.1-$(ARCH) modernprobe-builder"; \
	else \
		echo "Building modernprobe-builder image"; \
		docker build -t $(REPOSITORY)/modernprobe-builder:0.0.1-$(ARCH) modernprobe-builder; \
	fi

initrd-builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building initrd-builder image: docker build -t initrd-builder:0.0.1 initrd-builder"; \
	else \
		echo "Building initrd-builder image"; \
		docker build -t initrd-builder:0.0.1 initrd-builder; \
	fi

build-rootfs:
	@for rootfs_dir in $$(find . -type f -name "Dockerfile" -path "*$(ARCH)/*" -printf '%h\n'); do \
		version=$$(basename $$rootfs_dir); \
		distro=$$(basename $$(dirname $$rootfs_dir)); \
		image=$(REPOSITORY)/$$distro-image:$$version-$(ARCH); \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Building rootfs image: docker build -t $$image $$rootfs_dir"; \
		else \
			echo "Building rootfs image: $$image"; \
			docker build -t $$image $$rootfs_dir; \
		fi; \
	done

build-kernel: initrd-builder
	@for kernel_dir in $$(find . -type f -name "Dockerfile.kernel" -path "*$(ARCH)/*" -printf '%h\n'); do \
		version=$$(basename $$kernel_dir); \
		distro=$$(basename $$(dirname $$kernel_dir)); \
		image=$(REPOSITORY)/$$distro-kernel:$$version-$(ARCH); \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Building kernel image: docker build -t $$image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir"; \
		else \
			echo "Building kernel image: $$image"; \
			docker build -t $$image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir; \
		fi; \
	done

docker-push:
	@for image in $$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "$(REPOSITORY)" | grep "$(ARCH)"); do \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Pushing image: docker push $$image"; \
		else \
			echo "Pushing image: $$image"; \
			docker push $$image; \
		fi; \
	done

%: initrd-builder
	@version=$$(basename $@); \
		distro=$$(basename $$(dirname $@)); \
		arch=$(ARCH); \
		rootfs_dir=$$(find . -type d -path "./$$arch/$$distro/$$version"); \
		kernel_dir=$$(find . -type d -path "./$$arch/$$distro/$$version"); \
		rootfs_image=$(REPOSITORY)/$$distro-image:$$version-$$arch; \
		kernel_image=$(REPOSITORY)/$$distro-kernel:$$version-$$arch; \
		if [ -n "$$rootfs_dir" ]; then \
			if [ "$(DRY_RUN)" = "true" ]; then \
				echo "Dry run: Building rootfs image: docker build -t $$rootfs_image $$rootfs_dir"; \
			else \
				echo "Building rootfs image: $$rootfs_image"; \
				docker build -t $$rootfs_image $$rootfs_dir; \
			fi; \
		fi; \
		if [ -n "$$kernel_dir" ]; then \
			if [ "$(DRY_RUN)" = "true" ]; then \
				echo "Dry run: Building kernel image: docker build -t $$kernel_image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir"; \
			else \
				echo "Building kernel image: $$kernel_image"; \
				docker build -t $$kernel_image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir; \
			fi; \
		fi;

generate-yaml:
	@echo "machines:" > $(YAML_FILE)
	@for rootfs_image in $$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "$(REPOSITORY)" | grep "$(ARCH)" | grep "\-image" | sort); do \
		version=$$(echo $$rootfs_image | cut -d':' -f2 | cut -d'-' -f1); \
		distro=$$(echo $$rootfs_image | cut -d'/' -f2 | cut -d'-' -f1); \
		arch=$(ARCH); \
		kernel_image=$$(docker images --format "{{.Repository}}:{{.Tag}}" | grep "$(REPOSITORY)/$$distro-kernel:$$version-$$arch"); \
		if [ -n "$$kernel_image" ]; then \
			echo "  - {name: \"$$distro-$$version\", kernel: \"$$kernel_image\", rootfs: \"$$rootfs_image\", arch: \"$$arch\"}" >> $(YAML_FILE); \
		fi; \
	done

build-all: build-rootfs build-kernel
