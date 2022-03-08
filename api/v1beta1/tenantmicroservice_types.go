/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TenantMicroserviceSpec defines the desired state of TenantMicroservice
type TenantMicroserviceSpec struct {
	// Tenant-specific microservice configuration.
	Configuration EntityConfiguration `json:"configuration"`
}

// TenantMicroserviceStatus defines the observed state of TenantMicroservice
type TenantMicroserviceStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=dctm
//+kubebuilder:subresource:status

// TenantMicroservice is the Schema for the tenantmicroservices API
type TenantMicroservice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantMicroserviceSpec   `json:"spec,omitempty"`
	Status TenantMicroserviceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TenantMicroserviceList contains a list of TenantMicroservice
type TenantMicroserviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantMicroservice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TenantMicroservice{}, &TenantMicroserviceList{})
}
