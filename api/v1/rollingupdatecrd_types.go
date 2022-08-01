/*
Copyright 2022.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RollingUpdateCrdSpec defines the desired state of RollingUpdateCrd
type RollingUpdateCrdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RollingUpdateCrd. Edit rollingupdatecrd_types.go to remove/update
	DeploymentName string `json:"deploymentName,omitempty"`
}

// RollingUpdateCrdStatus defines the observed state of RollingUpdateCrd
type RollingUpdateCrdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RollingUpdateCrd is the Schema for the rollingupdatecrds API
type RollingUpdateCrd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RollingUpdateCrdSpec   `json:"spec,omitempty"`
	Status RollingUpdateCrdStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RollingUpdateCrdList contains a list of RollingUpdateCrd
type RollingUpdateCrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RollingUpdateCrd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RollingUpdateCrd{}, &RollingUpdateCrdList{})
}
