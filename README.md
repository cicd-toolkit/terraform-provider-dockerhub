=======
# Docker Hub Provider



![Build Status](https://github.com/cicd-toolkit/terraform-provider-dockerhub/actions/workflows/publish.yml/badge.svg)

Lifecycle management of Docker Hub using the v2 API.

This provider enables management of Docker Hub registries, groups, permissions and access tokens.

## Installation

The provider can be installed and managed automatically by Terraform. Sample `versions.tf` file:

```hcl
terraform {
  required_version = ">= 0.13"

  required_providers {
    dockerhub = {
      source  = "cicd-toolkit/dockerhub"
      version = ">= 0.0.15"
    }
  }
}
```

## Quick Start

```hcl
# Configure the Docker Hub Provider
provider "dockerhub" {
  username = "azurediamond"
  password = "hunter2"
}


data "dockerhub_tag" "example" {
  most_recent = true

  owners = ["self"]
}

```

## Development Guide

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make terraform-provider-dockerhub`. This will build the provider and put the provider binary in the local directory.

### Testing

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

### Installing

To install the provider locally (for instance for manually testing) run

```sh
$ make install
```

This installs the provider to `~/.terraform.d/`, removes `.terraform.lock.hcl` and runs `terraform init` in the local folder.
