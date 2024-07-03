# Deploying DutyController with Helm

DutyController is a Kubernetes Operator designed to facilitate the integration and management of incident management services like PagerDuty directly through Kubernetes resources. This guide covers the steps to deploy DutyController using Helm, a package manager for Kubernetes that simplifies the deployment and management of applications.

## Prerequisites

Before you begin, ensure you have the following prerequisites installed and configured:

- Kubernetes cluster
- Helm 3

## Adding the DutyController Helm Repository

To install DutyController, you first need to add the Helm repository that contains the DutyController chart. Run the following command to add the repository:

```bash
helm repo add duty https://mattgialelis.github.io/dutycontroller/
```

This command adds the DutyController repository under the name `duty`. Once added, you can search for the DutyController chart in the repository.

## Installing DutyController

With the repository added, you can now install DutyController into your Kubernetes cluster. Use the following command to install DutyController using Helm:

```bash
helm install dutycontroller duty/dutycontroller
```

This command deploys DutyController on the Kubernetes cluster in the default configuration. The `helm install` command deploys a new instance of DutyController and assigns it the release name `dutycontroller`.

## Verifying the Installation

After installing, you can verify that DutyController is running by checking the deployed pods:

```bash
kubectl get pods
```

Look for the pods that start with `dutycontroller-` prefix. If the pods are running, DutyController has been successfully deployed.

## Updating DutyController

To update DutyController to the latest version, first update your Helm repository to fetch the latest charts:

```bash
helm repo update
```

Then, upgrade DutyController using Helm:

```bash
helm upgrade dutycontroller duty/dutycontroller
```

This command upgrades DutyController to the latest version available in the repository.

## Uninstalling DutyController

If you wish to remove DutyController from your cluster, you can uninstall it using Helm:

```bash
helm uninstall dutycontroller
```

This command removes all the Kubernetes components associated with the release and deletes the `dutycontroller` release.

