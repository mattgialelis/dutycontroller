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

// BusinessServiceSpec defines the desired state of BusinessService
type BusinessServiceSpec struct {
	// Remember: Run "make" to regenerate code after modifying this file

	// Name is the name of the BusinessService we want to create
	Description    string `json:"description"`
	PointOfContact string `json:"pointOfContact"`
	Team           string `json:"team"`
}

// BusinessServiceStatus defines the observed state of BusinessService
type BusinessServiceStatus struct {
	ID         string             `json:"id,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// BusinessService is the Schema for the businessservices API
type BusinessService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BusinessServiceSpec   `json:"spec,omitempty"`
	Status BusinessServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BusinessServiceList contains a list of BusinessService
type BusinessServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BusinessService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BusinessService{}, &BusinessServiceList{})
}
