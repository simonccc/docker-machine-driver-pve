# Sample Ubuntu Server template

This is sample Packer configuration for creating an Ubuntu Server template compatible with the [Docker/Rancher Machine driver](../../../README.md). This template is intended **for testing purposes only**. It has **NOT** been security-hardened or validated for production deployments.

## Requirements

To avoid managing passwords, this template uses SSH pubkey authentication when connecting to the machine. To achieve that:

* SSH public key must be configured in `variables.pkr.hcl` (see [usage](#usage)),
* corresponding SSH identity must be present in the local SSH agent when building the template.

## Usage

1. Initialize Packer plugins:

    ```sh
    packer init ./ubuntu-server.pkr.hcl
    ```

1. Create a variables file:

    ```sh
    cp ./variables.example.pkr.hcl ./variables.pkr.hcl
    ```

1. Fill out `variables.pkr.hcl` file with values for your environment.

1. Ensure SSH identity is added to the SSH Agent (see [requirements](#requirements)):

    ```sh
    ssh-add -l
    ```

1. Run a build:

    ```sh
    packer build --var-file ./variables.pkr.hcl ./ubuntu-server.pkr.hcl
    ```
