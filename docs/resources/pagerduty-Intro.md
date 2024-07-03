# Kubernetes CRD for PagerDuty Integration

The Custom Resource Definition (CRD) provided by this project is designed to facilitate the seamless onboarding of services into PagerDuty within a Kubernetes-native framework. This integration aims to leverage the power and flexibility of Kubernetes to manage and automate the configuration of incident management services, specifically tailored for PagerDuty.

## Features and Enhancements

- **Kubernetes-Native Approach:** The CRD aligns with Kubernetes principles, allowing users to manage PagerDuty services as part of their Kubernetes resource configurations. This integration ensures that the setup and maintenance of PagerDuty services feel natural to Kubernetes users, promoting infrastructure as code practices.

- **Close Alignment with PagerDuty API:** All CRDs are implemented to closely mirror the functionality provided by the [PagerDuty API](https://developer.pagerduty.com/api-reference/), ensuring that users have access to the comprehensive features offered by PagerDuty directly from their Kubernetes environment.

- **Enhanced Functionality:** While maintaining close alignment with the PagerDuty API, the CRD introduces minor improvements to streamline operations. Notably, it simplifies the creation and linkage of business services directly through a single service resource. This enhancement reduces the complexity and overhead associated with managing interconnected services, making it easier to set up comprehensive incident response workflows.



## Environment Variables Needed for PagerDuty

To ensure the controller can securely interact with PagerDuty's API, it's necessary to set up certain environment variables. Below is a table detailing the required environment variable:

| Environment Variable | Description                                                                        | Documentation Link                                                                                                                              |
| -------------------- | ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `PAGERDUTY_TOKEN`    | This token is used to authenticate the controller's requests to the PagerDuty API. | [Generating a General Access REST API Key](https://support.pagerduty.com/docs/api-access-keys#section-generating-a-general-access-rest-api-key) |

### Setting Up `PAGERDUTY_TOKEN`

To interact with PagerDuty's API, the controller requires a REST API key. This key is specified as an environment variable (`PAGERDUTY_TOKEN`) that the controller reads at runtime. Follow the instructions in the provided documentation link to generate a General Access REST API key in PagerDuty. Once generated, ensure this key is securely stored and provided to the controller as an environment variable.



