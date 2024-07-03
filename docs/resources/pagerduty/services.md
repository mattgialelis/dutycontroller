# Services Custom Resource Definition (CRD)

The Services CRD allows for the declarative configuration of PagerDuty services within a Kubernetes environment. This document outlines how to create and manage these resources, providing a Kubernetes-native approach to incident management.

## Example Resource

Below is an example of a Services resource definition:

```yaml
apiVersion: pagerduty.dutycontroller.io/v1beta1
kind: Services
metadata:
  name: my-service
  labels:
    app.kubernetes.io/name: dutycontroller
    app.kubernetes.io/managed-by: kustomize
spec:
  acknowledgeTimeout: 300
  autoResolveTimeout: 600
  businessService: "businessservice-sample"
  description: "My services"
  escalationPolicy: "test1-ep"
  status: "Active"
```



## Felid Descriptions

Each field in the Services CRD plays a crucial role in defining how the service interacts with PagerDuty. The table below provides detailed descriptions of these fields:

| Field                | Description                                                                                                                                                                                                   | Type    | Required |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- | -------- |
| `description`        | A description of the PagerDuty service.                                                                                                                                                                       | string  | Yes      |
| `status`             | The status of the PagerDuty service ("Active" or "Disabled").                                                                                                                                                 | string  | Yes      |
| `escalationPolicy`   | The name of the escalation policy associated with the service. The corresponding escalation policy ID is looked up in PagerDuty.                                                                              | string  | Yes      |
| `autoResolveTimeout` | (Optional) The number of seconds before an incident is automatically resolved.                                                                                                                                | integer | No       |
| `acknowledgeTimeout` | (Optional) The number of seconds before an unacknowledged incident escalates.                                                                                                                                 | integer | No       |
| `businessService`    | The name of the business service associated with the service. The corresponding business service ID is first looked up within the same Kubernetes namespace. If not found, it is then looked up in PagerDuty. | string  | No       |



## Additional Notes

- **Finalizer Integration:** The controller leverages finalizers to ensure the orderly deletion of PagerDuty services, aligning the lifecycle of Kubernetes resources with their PagerDuty counterparts. This approach guarantees that resources are not left in an inconsistent state upon deletion.
- **Idempotency and Duplication Prevention:** In cases where a specified service already exists within PagerDuty, the controller opts not to create a new instance. Instead, it records a log entry indicating the existence of the service. This behavior is crucial for maintaining idempotency and avoiding the unintended proliferation of duplicate services.
