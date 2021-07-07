/*
Copyright 2021.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TaintSpec defines the desired state of Taint
type TaintSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// NodeSelector is a label selector to point Node you want to add taints
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Taints are node taints you want to add to Nodes matched with NodeSelector
	// +kubebuilder:validation:MinItems=1
	Taints []corev1.Taint `json:"taints,omitempty"`
}

// TaintStatus defines the observed state of Taint
type TaintStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Taint is the Schema for the taints API
type Taint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TaintSpec   `json:"spec,omitempty"`
	Status TaintStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TaintList contains a list of Taint
type TaintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Taint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Taint{}, &TaintList{})
}
