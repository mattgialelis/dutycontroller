# Orchestration Routes CRD Documentation

This document provides detailed guidance on configuring Orchestration Routes using the Custom Resource Definition (CRD) in a Kubernetes environment for efficient event routing to PagerDuty services.

## Example Resource

```yaml
apiVersion: pagerduty.dutycontroller.io/v1beta1
kind: Orchestrationroutes
metadata:
    labels:
        app.kubernetes.io/name: dutycontroller
        app.kubernetes.io/managed-by: kustomize
    name: orchestrationroutes-sample
spec:
    serviceRoutes:
        - eventOrchestration: "Grafana AlertManager"
            label: "My AlertManager"
            serviceRef: "test"
            conditions:
                - "event.custom_details.component matches part 'my-service-b'"
        - eventOrchestration: "Grafana AlertManager"
            label: "My Service AlertManager"
            serviceRef: "my-service"
            conditions:
                - "event.custom_details.component matches part 'my-service-deleteme'"
```

## Field Descriptions

The table below provides a comprehensive overview of the fields within the Orchestration Routes CRD, detailing their purpose, data type, and whether they are mandatory.

| Field                | Description                                                                                                                                                                                                         | Type     | Required |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | -------- |
| `eventOrchestration` | Specifies the event orchestration integration in PagerDuty, such as "Grafana AlertManager". This field identifies the external system from which events are received.                                               | string   | Yes      |
| `label`              | Assigns a unique label to the routing rule for easy identification within PagerDuty.                                                                                                                                | string   | Yes      |
| `serviceRef`         | References the `Services` custom resource that defines the PagerDuty service targeted for event routing. The system searches for a matching `Services` resource within the same namespace or directly in PagerDuty. | string   | Yes      |
| `conditions`         | Lists JMESPath expressions that filter events. All conditions must evaluate to `true` for an event to be routed to the specified service. This allows for precise control over which events trigger alerts.         | []string | No       |

## Additional Notes

- **Flexibility in Event Routing:** The Orchestration Routes CRD offers granular control over event routing, enabling users to define complex routing logic based on the content of the event, thereby ensuring that alerts are directed to the most appropriate service or team.
- **Integration with External Systems:** By specifying `eventOrchestration`, users can seamlessly integrate with a variety of external monitoring and alerting systems, enhancing the operational efficiency of incident management workflows.
