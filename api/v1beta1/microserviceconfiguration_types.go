/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MicroserviceConfigurationSpec defines the desired state of MicroserviceConfiguration
type MicroserviceConfigurationSpec struct {
	// Instance configuration information.
	Configuration EntityConfiguration `json:"configuration"`
}

// MicroserviceConfigurationStatus defines the observed state of MicroserviceConfiguration
type MicroserviceConfigurationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=dcmc
//+kubebuilder:subresource:status

// MicroserviceConfiguration is the Schema for the microserviceconfigurations API
type MicroserviceConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MicroserviceConfigurationSpec   `json:"spec,omitempty"`
	Status MicroserviceConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MicroserviceConfigurationList contains a list of MicroserviceConfiguration
type MicroserviceConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicroserviceConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MicroserviceConfiguration{}, &MicroserviceConfigurationList{})
}
