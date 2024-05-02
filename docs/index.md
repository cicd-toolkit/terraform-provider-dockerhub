---
layout: ""
page_title: "Provider: dockerhub"
description: |-
  Register hub.docker.com repositories.
---

# Docker Hub Provider

Lifecycle management of Docker Hub using the v2 API.

This provider enables management of Docker Hub registries, groups, permissions and access tokens.

## Installation

The provider can be installed and managed automatically by Terraform. Sample `versions.tf` file:

```terraform
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

## Configuration

```terraform
provider "dockerhub" {
  # Note: This cannot be a Personal Access Token
  username = "USERNAME" # or use DOCKER_USERNAME environment variable
  password = "PASSWORD" # or use DOCKER_PASSWORD environment variable
}
```

### Argument Reference

- **password** (String) Password for authentication.
- **username** (String) Username for authentication.

## Example

```terraform
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
