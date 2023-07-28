# Images

Makefile present in this directory is specifically designed to generate the static Docker images required by Ignite to run tests on different Linux distributions. The workflow provided by this Makefile is designed to be straightforward, consisting of three main commands:

1. `build-all`: This target builds all the necessary Docker images for the different versions and distributions required for testing with Firecracker.

2. `docker-push`: Optionally, you can use this target to push the resulting Docker images to a Docker Hub registry for easier distribution and access.

3. `generate-yaml`: This target allows you to generate a YAML file (`images.yaml`) containing the matrix of new image information. The generated YAML file can be conveniently copied to the variables file of Ansible to keep the test environment up to date.

## Prerequisites

Before using the Makefile, ensure you have the following installed:

- Docker: The containerization platform used for building and pushing images.

## Makefile Targets

The Makefile provides several targets, each serving a specific purpose:

- `initrd-builder`: This target builds the `initrd-builder` image, necessary for creating the initrd for each image.

- `builder`, `modernprobe-builder`: These targets build specific builder images used by the CI system to prepare all the precompiled files for the tests.

- `build-rootfs` and `build-kernel`: These targets build root filesystem and kernel images, respectively. The `build-kernel` target depends on `initrd-builder`, which must be built first.

- `docker-push`: This target pushes the built images to a Docker Hub registry. You can use this step to make the images accessible to other systems.

- `generate-yaml`: This target generates a YAML file named `images.yaml`, which contains information about the built images. The YAML file includes details about the kernel and rootfs images for each version and distribution. This generated YAML file can be conveniently copied to the variables file of Ansible to keep the test environment up to date.

- `build-all`: This target is a convenience target that sequentially builds both root filesystem and kernel images.

## Usage

The typical workflow for using this Makefile is as follows:

1. Build the `initrd-builder` image first, which is required for creating the initrd for each image:

```
make initrd-builder
```

2. Build the specific builder images (`builder`, `modernprobe-builder`) used by the CI system:

```
make builder
make modernprobe-builder
```

3. Build all the required images for testing with Firecracker using the following command:

```
make build-all
```

4. Optionally, push the built images to a Docker Hub registry with:

```
make docker-push
```

5. Generate the YAML file containing the image matrix with:

```
make generate-yaml
```

After running these commands, you will have the necessary Docker images for your Firecracker test environment, and the image matrix will be available in the `images.yaml` file. You can then easily integrate this information into your Ansible setup.

## Customization

You can customize the Makefile to suit your specific requirements. The variables you can modify include:

- `DRY_RUN`: Set this variable to `true` for a dry run, where the build commands will be printed but not executed.

- `REPOSITORY`: The Docker repository where the built images will be tagged and pushed.

- `ARCH`: The architecture for which the images will be built. By default, it will use the output of `uname -p`.

- `YAML_FILE`: The name of the YAML file that will be generated by the `generate-yaml` target.

Feel free to adjust these variables to match your desired configuration.