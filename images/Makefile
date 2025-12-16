DRY_RUN ?= false
PUSH ?= false
LATEST ?= false
TAG ?= main
CLEANUP ?= false
REPOSITORY ?= ghcr.io/falcosecurity/kernel-testing
ARCH ?= $(shell uname -m)
YAML_FILE ?= images.yaml
CMAKE_VERSION ?= $(shell cat CMAKE_VERSION)

.PHONY: build-rootfs build-kernel generate-yaml build-all initrd-builder modernprobe-builder builder

builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building builder image: docker build -t $(REPOSITORY)/builder:$(ARCH)-$(TAG) --build-arg CMAKE_VERSION=\"$(CMAKE_VERSION)\" builder"; \
	else \
		echo "Building builder image"; \
		image=$(REPOSITORY)/builder:$(ARCH); \
		docker build \
			-t $$image-$(TAG) \
			--build-arg CMAKE_VERSION="$(CMAKE_VERSION)" \
			 builder || echo "FAIL: $$image-$(TAG)" >> failed.txt; \
		if [ "$(PUSH)" = "true" ]; then \
			echo "Pushing image: $$image-$(TAG)"; \
			docker push $$image-$(TAG); \
			if [ "$(LATEST)" = "true" ]; then \
				echo "Pushing image: $$image-latest"; \
				docker tag $$image-$(TAG) $$image-latest; \
				docker push $$image-latest; \
			fi; \
		fi; \
	fi

modernprobe-builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building modernprobe-builder image: docker build -t $(REPOSITORY)/modernprobe-builder:$(ARCH)-$(TAG) --build-arg CMAKE_VERSION=\"$(CMAKE_VERSION)\" modernprobe-builder"; \
	else \
		echo "Building modernprobe-builder image"; \
		image=$(REPOSITORY)/modernprobe-builder:$(ARCH); \
		docker build \
			-t $$image-$(TAG) \
			--build-arg CMAKE_VERSION="$(CMAKE_VERSION)" \
			modernprobe-builder || echo "FAIL: $$image-$(TAG)" >> failed.txt; \
		if [ "$(PUSH)" = "true" ]; then \
			echo "Pushing image: $$image-$(TAG)"; \
			docker push $$image-$(TAG); \
			if [ "$(LATEST)" = "true" ]; then \
				echo "Pushing image: $$image-latest"; \
				docker tag $$image-$(TAG) $$image-latest; \
				docker push $$image-latest; \
			fi; \
		fi; \
	fi

# No need to push intrd-builder image; it is just used locally as base to other images.
initrd-builder:
	if [ "$(DRY_RUN)" = "true" ]; then \
		echo "Dry run: Building initrd-builder image: docker build -t initrd-builder:0.0.1 initrd-builder"; \
	else \
		echo "Building initrd-builder image"; \
		docker build -t initrd-builder:0.0.1 initrd-builder || echo "FAIL: initrd-builder" >> failed.txt; \
	fi

build-rootfs:
	@for rootfs_dir in $$(find . -type f -name "Dockerfile" -path "*$(ARCH)/*" -printf '%h\n'); do \
		version=$$(basename $$rootfs_dir); \
		distro=$$(basename $$(dirname $$rootfs_dir)); \
		image=$(REPOSITORY)/$$distro-image:$$version-$(ARCH); \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Building rootfs image: docker build -t $$image-$(TAG) --build-arg CMAKE_VERSION=\"$(CMAKE_VERSION)\" $$rootfs_dir"; \
		else \
			echo "Building rootfs image: $$image-$(TAG)"; \
			docker build \
				-t $$image-$(TAG) \
				--build-arg CMAKE_VERSION="$(CMAKE_VERSION)" \
				$$rootfs_dir || echo "FAIL: $$image-$(TAG)" >> failed.txt; \
			if [ "$(PUSH)" = "true" ]; then \
				echo "Pushing image: $$image-$(TAG)"; \
				docker push $$image-$(TAG); \
				if [ "$(LATEST)" = "true" ]; then \
					echo "Pushing image: $$image-latest"; \
					docker tag $$image-$(TAG) $$image-latest; \
					docker push $$image-latest; \
				fi; \
			fi; \
			if [ "$(CLEANUP)" = "true" ]; then \
				echo "Cleaning image: $$image-$(TAG)"; \
				docker image rm -f $$image-$(TAG); \
				docker image rm -f $$image-latest; \
				docker system prune -f; \
			fi; \
		fi; \
	done

build-kernel: initrd-builder
	@for kernel_dir in $$(find . -type f -name "Dockerfile.kernel" -path "*$(ARCH)/*" -printf '%h\n'); do \
		version=$$(basename $$kernel_dir); \
		distro=$$(basename $$(dirname $$kernel_dir)); \
		image=$(REPOSITORY)/$$distro-kernel:$$version-$(ARCH); \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Building kernel image: docker build -t $$image-$(TAG) -f $$kernel_dir/Dockerfile.kernel $$kernel_dir"; \
		else \
			echo "Building kernel image: $$image-$(TAG)"; \
			docker build -t $$image-$(TAG) -f $$kernel_dir/Dockerfile.kernel $$kernel_dir || echo "FAIL: $$image-$(TAG)" >> failed.txt; \
			if [ "$(PUSH)" = "true" ]; then \
				echo "Pushing image: $$image-$(TAG)"; \
				docker push $$image-$(TAG); \
				if [ "$(LATEST)" = "true" ]; then \
					echo "Pushing image: $$image-latest"; \
					docker tag $$image-$(TAG) $$image-latest; \
					docker push $$image-latest; \
				fi; \
			fi; \
			if [ "$(CLEANUP)" = "true" ]; then \
				echo "Cleaning image: $$image-$(TAG)"; \
				docker image rm -f $$image-$(TAG); \
				docker image rm -f $$image-latest; \
				docker system prune -f; \
			fi; \
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
			echo "Dry run: Building rootfs image: docker build -t $$rootfs_image --build-arg CMAKE_VERSION=\"$(CMAKE_VERSION)\" $$rootfs_dir"; \
		else \
			echo "Building rootfs image: $$rootfs_image"; \
			docker build \
				-t $$rootfs_image \
				--build-arg CMAKE_VERSION="$(CMAKE_VERSION)" \
				$$rootfs_dir; \
			if [ "$(PUSH)" = "true" ]; then \
				echo "Pushing image: $$rootfs_image"; \
				docker push $$rootfs_image; \
			fi; \
		fi; \
	fi; \
	if [ -n "$$kernel_dir" ]; then \
		if [ "$(DRY_RUN)" = "true" ]; then \
			echo "Dry run: Building kernel image: docker build -t $$kernel_image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir"; \
		else \
			echo "Building kernel image: $$kernel_image"; \
			docker build -t $$kernel_image -f $$kernel_dir/Dockerfile.kernel $$kernel_dir; \
			if [ "$(PUSH)" = "true" ]; then \
				echo "Pushing image: $$kernel_image"; \
				docker push $$kernel_image; \
			fi; \
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

build-all: build-kernel build-rootfs modernprobe-builder builder
