/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServicesSpec defines the desired state of Services
type ServicesSpec struct {
	Description        string `json:"description"`
	Status             string `json:"status"`
	EscalationPolicy   string `json:"escalationPolicy"`
	AutoResolveTimeout int    `json:"autoResolveTimeout"`
	AcknowedgeTimeout  int    `json:"acknowledgeTimeout"`
	BusinessService    string `json:"businessService"`
}

// ServicesStatus defines the observed state of Services
type ServicesStatus struct {
	ID string `json:"id,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Services is the Schema for the services API
type Services struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServicesSpec   `json:"spec,omitempty"`
	Status ServicesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServicesList contains a list of Services
type ServicesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Services `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Services{}, &ServicesList{})
}
