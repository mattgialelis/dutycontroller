## Duty Controller Docs Site

DutyController is a Kubernetes Operator crafted to streamline the integration and management of incident management services, such as PagerDuty, directly through Kubernetes resources. This project bridges the gap in Kubernetes' native capabilities, providing a unified approach to managing incident response configurations seamlessly within a Kubernetes ecosystem.

For more details on integrating and managing PagerDuty services, refer to the [Resources section](#resources).

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
