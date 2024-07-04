# Business Service CRD Documentation

This document provides a comprehensive guide on the Business Service Custom Resource Definition (CRD) for integrating with PagerDuty services in a Kubernetes environment.

## Example Resource

Below is an example of how to define a Business Service resource:

```yaml
apiVersion: pagerduty.dutycontroller.io/v1beta1
kind: BusinessService
metadata:
  labels:
    app.kubernetes.io/name: dutycontroller
  name: businessservice-sample
spec:
  description: "Example DescriptionNew"
  pointOfContact: "Example Contact"
  team: "Example Team"
```


## Feild Descriptions

The table below details the fields available in the Business Service CRD, their descriptions, types, and whether they are required.

| Field            | Description                                                                                                                             | Type   | Required |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------- | ------ | -------- |
| `description`    | A description of the business service.                                                                                                  | string | Yes      |
| `pointOfContact` | The primary point of contact (person or team name) for the business service.                                                            | string | Yes      |
| `team`           | The name of the PagerDuty team associated with the business service. The corresponding team ID is looked up automatically in PagerDuty. | string | Yes      |


## Additional Notes

- **Finalizer Integration:** The controller leverages finalizers to ensure the orderly deletion of PagerDuty business services, aligning the lifecycle of Kubernetes resources with their PagerDuty counterparts. This approach guarantees that resources are not left in an inconsistent state upon deletion.
- **Idempotency and Duplication Prevention:** In cases where a specified business service already exists within PagerDuty, the controller opts not to create a new instance. Instead, it records a log entry indicating the existence of the service. This behavior is crucial for maintaining idempotency and avoiding the unintended proliferation of duplicate services.
