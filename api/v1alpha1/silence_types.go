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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SilenceSpec defines the desired state of Silence
type SilenceSpec struct {
	AlertmanagerSilence `json:",inline"`
}

// SilenceStatus defines the observed state of Silence
type SilenceStatus struct {
	AlertmanagerReferences []AlertmanagerReference `json:"alertmanagerReference,omitempty"`
	Conditions             []metav1.Condition      `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Silence is the Schema for the silences API
type Silence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SilenceSpec   `json:"spec,omitempty"`
	Status SilenceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SilenceList contains a list of Silence
type SilenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Silence `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Silence{}, &SilenceList{})
}
