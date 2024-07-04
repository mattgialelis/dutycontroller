# DutyController Helm Chart Configuration

This document outlines the configurable values for the DutyController Helm chart. These values can be adjusted to tailor the deployment of DutyController to your Kubernetes cluster.

## Configuration Values

### General Configuration

- `nameOverride`: Overrides the chart name for resources (default: `""`).
- `fullnameOverride`: Overrides the full name of the resources (default: `""`).
- `imagePullSecrets`: Specifies docker-registry secret names as an array (default: `[]`).

### Controller Options

- `leaderElect`: Enables leader election for controller manager (default: `false`).
- `livenessProbe`: Configures the liveness probe for the controller (default: HTTP GET `/healthz`).
- `readinessProbe`: Configures the readiness probe for the controller (default: HTTP GET `/readyz`).
- `replicaCount`: Number of DutyController replicas (default: `1`).

### Image Configuration

- `image.repository`: The repository for the DutyController image (default: `ghcr.io/mattgialelis/dutycontroller`).
- `image.pullPolicy`: Image pull policy (default: `IfNotPresent`).
- `image.tag`: The image tag to use (default: chart `appVersion`).

### Resource Limits and Requests

- `resources`: CPU/memory resource requests/limits (default: `100m` CPU, `128Mi` memory).

### Service Account Configuration

- `serviceAccount.create`: Specifies whether a service account should be created (default: `true`).
- `serviceAccount.automount`: Automount service account token (default: `true`).
- `serviceAccount.annotations`: Annotations to add to the service account (default: `{}`).
- `serviceAccount.name`: The name of the service account to use.

### RBAC Configuration

- `rbac.create`: Specifies whether RBAC resources should be created (default: `true`).

### Monitoring Configuration

- `monitoring.enabled`: Enable metrics port (default: `false`).
- `monitoring.port`: Port for metrics endpoint (default: `8080`).
- `monitoring.serviceMonitor.enabled`: Enable service monitor creation (default: `false`).
- `monitoring.serviceMonitor.additionalLabels`: Additional labels for service monitor (default: `{}`).

### Service Configuration

- `service.type`: Type of service to create (default: `ClusterIP`).
- `service.ports`: Ports configuration for the service.

### Additional Configuration

- `extraEnv`: Extra environment variables to add to the container.
- `envFrom`: Secrets that are mounted as `envFrom` on the pod.
- `podAnnotations`: Annotations to add to the pods.
- `podLabels`: Labels to add to the pods.
- `nodeSelector`: Node labels for pod assignment.
- `tolerations`: Tolerations for pod assignment.
- `affinity`: Affinity for pod assignment.
- `extraObjects`: Create dynamic manifests via values.

## Example Usage

To customize the deployment, create a `values.yaml` file with your desired values and pass it during installation:

```bash
helm install dutycontroller duty/dutycontroller -f values.yaml
```


## Example Values File

```yaml
image:
  repository: ghcr.io/mattgialelis/dutycontroller

envFrom:
  - secretRef:
      name: dutycontroller-envs

extraObjects:
  - apiVersion: external-secrets.io/v1beta1
    kind: ExternalSecret
    metadata:
      name: dutycontroller-envs
    spec:
      data:
      - remoteRef:
          conversionStrategy: Default
          decodingStrategy: None
          key: pagerduty-token
        secretKey: PAGERDUTY_TOKEN
      refreshInterval: 1h
      secretStoreRef:
        kind: ClusterSecretStore
        name: myClusterStore
```
