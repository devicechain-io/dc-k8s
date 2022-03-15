//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"encoding/json"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityConfiguration) DeepCopyInto(out *EntityConfiguration) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityConfiguration.
func (in *EntityConfiguration) DeepCopy() *EntityConfiguration {
	if in == nil {
		return nil
	}
	out := new(EntityConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfiguration) DeepCopyInto(out *InstanceConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfiguration.
func (in *InstanceConfiguration) DeepCopy() *InstanceConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationList) DeepCopyInto(out *InstanceConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationList.
func (in *InstanceConfigurationList) DeepCopy() *InstanceConfigurationList {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationSpec) DeepCopyInto(out *InstanceConfigurationSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationSpec.
func (in *InstanceConfigurationSpec) DeepCopy() *InstanceConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationStatus) DeepCopyInto(out *InstanceConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationStatus.
func (in *InstanceConfigurationStatus) DeepCopy() *InstanceConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCreateRequest) DeepCopyInto(out *InstanceCreateRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCreateRequest.
func (in *InstanceCreateRequest) DeepCopy() *InstanceCreateRequest {
	if in == nil {
		return nil
	}
	out := new(InstanceCreateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGetRequest) DeepCopyInto(out *InstanceGetRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGetRequest.
func (in *InstanceGetRequest) DeepCopy() *InstanceGetRequest {
	if in == nil {
		return nil
	}
	out := new(InstanceGetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroservicGetRequest) DeepCopyInto(out *MicroservicGetRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroservicGetRequest.
func (in *MicroservicGetRequest) DeepCopy() *MicroservicGetRequest {
	if in == nil {
		return nil
	}
	out := new(MicroservicGetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Microservice) DeepCopyInto(out *Microservice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Microservice.
func (in *Microservice) DeepCopy() *Microservice {
	if in == nil {
		return nil
	}
	out := new(Microservice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Microservice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceConfiguration) DeepCopyInto(out *MicroserviceConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceConfiguration.
func (in *MicroserviceConfiguration) DeepCopy() *MicroserviceConfiguration {
	if in == nil {
		return nil
	}
	out := new(MicroserviceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicroserviceConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceConfigurationGetRequest) DeepCopyInto(out *MicroserviceConfigurationGetRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceConfigurationGetRequest.
func (in *MicroserviceConfigurationGetRequest) DeepCopy() *MicroserviceConfigurationGetRequest {
	if in == nil {
		return nil
	}
	out := new(MicroserviceConfigurationGetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceConfigurationList) DeepCopyInto(out *MicroserviceConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MicroserviceConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceConfigurationList.
func (in *MicroserviceConfigurationList) DeepCopy() *MicroserviceConfigurationList {
	if in == nil {
		return nil
	}
	out := new(MicroserviceConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicroserviceConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceConfigurationSpec) DeepCopyInto(out *MicroserviceConfigurationSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceConfigurationSpec.
func (in *MicroserviceConfigurationSpec) DeepCopy() *MicroserviceConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(MicroserviceConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceConfigurationStatus) DeepCopyInto(out *MicroserviceConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceConfigurationStatus.
func (in *MicroserviceConfigurationStatus) DeepCopy() *MicroserviceConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(MicroserviceConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceCreateRequest) DeepCopyInto(out *MicroserviceCreateRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceCreateRequest.
func (in *MicroserviceCreateRequest) DeepCopy() *MicroserviceCreateRequest {
	if in == nil {
		return nil
	}
	out := new(MicroserviceCreateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceList) DeepCopyInto(out *MicroserviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Microservice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceList.
func (in *MicroserviceList) DeepCopy() *MicroserviceList {
	if in == nil {
		return nil
	}
	out := new(MicroserviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MicroserviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceListRequest) DeepCopyInto(out *MicroserviceListRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceListRequest.
func (in *MicroserviceListRequest) DeepCopy() *MicroserviceListRequest {
	if in == nil {
		return nil
	}
	out := new(MicroserviceListRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceSpec) DeepCopyInto(out *MicroserviceSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceSpec.
func (in *MicroserviceSpec) DeepCopy() *MicroserviceSpec {
	if in == nil {
		return nil
	}
	out := new(MicroserviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicroserviceStatus) DeepCopyInto(out *MicroserviceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroserviceStatus.
func (in *MicroserviceStatus) DeepCopy() *MicroserviceStatus {
	if in == nil {
		return nil
	}
	out := new(MicroserviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantCreateRequest) DeepCopyInto(out *TenantCreateRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantCreateRequest.
func (in *TenantCreateRequest) DeepCopy() *TenantCreateRequest {
	if in == nil {
		return nil
	}
	out := new(TenantCreateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantGetRequest) DeepCopyInto(out *TenantGetRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantGetRequest.
func (in *TenantGetRequest) DeepCopy() *TenantGetRequest {
	if in == nil {
		return nil
	}
	out := new(TenantGetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroservice) DeepCopyInto(out *TenantMicroservice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroservice.
func (in *TenantMicroservice) DeepCopy() *TenantMicroservice {
	if in == nil {
		return nil
	}
	out := new(TenantMicroservice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantMicroservice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceByTenantRequest) DeepCopyInto(out *TenantMicroserviceByTenantRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceByTenantRequest.
func (in *TenantMicroserviceByTenantRequest) DeepCopy() *TenantMicroserviceByTenantRequest {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceByTenantRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceCreateRequest) DeepCopyInto(out *TenantMicroserviceCreateRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceCreateRequest.
func (in *TenantMicroserviceCreateRequest) DeepCopy() *TenantMicroserviceCreateRequest {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceCreateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceGetRequest) DeepCopyInto(out *TenantMicroserviceGetRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceGetRequest.
func (in *TenantMicroserviceGetRequest) DeepCopy() *TenantMicroserviceGetRequest {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceGetRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceList) DeepCopyInto(out *TenantMicroserviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantMicroservice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceList.
func (in *TenantMicroserviceList) DeepCopy() *TenantMicroserviceList {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantMicroserviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceSpec) DeepCopyInto(out *TenantMicroserviceSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceSpec.
func (in *TenantMicroserviceSpec) DeepCopy() *TenantMicroserviceSpec {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantMicroserviceStatus) DeepCopyInto(out *TenantMicroserviceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantMicroserviceStatus.
func (in *TenantMicroserviceStatus) DeepCopy() *TenantMicroserviceStatus {
	if in == nil {
		return nil
	}
	out := new(TenantMicroserviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantStatus) DeepCopyInto(out *TenantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantStatus.
func (in *TenantStatus) DeepCopy() *TenantStatus {
	if in == nil {
		return nil
	}
	out := new(TenantStatus)
	in.DeepCopyInto(out)
	return out
}
