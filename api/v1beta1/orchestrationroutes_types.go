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

// OrchestrationroutesSpec defines the desired state of Orchestrationroutes
type OrchestrationroutesSpec struct {
	EventOrchestrationName string   `json:"eventOrchestrationName"`
	Expression             []string `json:"expression"`
	ServiceName            string   `json:"routeTo"`
}

// OrchestrationroutesStatus defines the observed state of Orchestrationroutes
type OrchestrationroutesStatus struct {
	RouteServiceId string `json:"routeTo"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Orchestrationroutes is the Schema for the orchestrationroutes API
type Orchestrationroutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrchestrationroutesSpec   `json:"spec,omitempty"`
	Status OrchestrationroutesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OrchestrationroutesList contains a list of Orchestrationroutes
type OrchestrationroutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Orchestrationroutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Orchestrationroutes{}, &OrchestrationroutesList{})
}
