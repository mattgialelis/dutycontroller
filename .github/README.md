# DutyController

DutyController is a Kubernetes Operator crafted to streamline the integration and management of incident management services, such as PagerDuty, directly through Kubernetes resources. This project bridges the gap in Kubernetes' native capabilities, providing a unified approach to managing incident response configurations seamlessly within a Kubernetes ecosystem.


## Documentation

For a comprehensive guide and detailed instructions on how to use DutyController, please visit our [documentation site](https://mattgialelis.github.io/dutycontroller/).


## Getting Started

### Docker Image

The DutyController Docker image is hosted on GitHub Container Registry. You can pull the image using the following command:

```bash
docker pull ghcr.io/mattgialelis/dutycontroller/dutycontroller:latest
```



### Helm Chart

To simplify the deployment of DutyController in your Kubernetes cluster, we provide a Helm chart. Add the DutyController Helm repository and install the chart with:

```bash

helm repo add duty https://mattgialelis.github.io/dutycontroller/
helm install dutycontroller duty/dutycontroller

```

## How to Contribute

We welcome contributions from the community! Whether it's submitting a bug report, proposing a new feature, or contributing code, we encourage you to get involved. Please check out our CONTRIBUTING guide for more information on how to start contributing.
