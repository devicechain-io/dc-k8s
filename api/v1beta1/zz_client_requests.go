/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

// ------------------
// Instance Mangement
// ------------------

// Information required to create a DeviceChain instance.
type InstanceCreateRequest struct {
	Id              string
	Name            string
	Description     string
	ConfigurationId string
}

// Information required to get a DeviceChain instance.
type InstanceGetRequest struct {
	Id string
}

// ------------------
// Tenant Mangement
// ------------------

// Information required to create a DeviceChain tenant.
type TenantCreateRequest struct {
	InstanceId  string
	TenantId    string
	Name        string
	Description string
}

// Information required to get a tenant.
type TenantGetRequest struct {
	InstanceId string
	TenantId   string
}

// ----------------------
// Microservice Mangement
// ----------------------

// Information required to get a microservice configuration.
type MicroserviceConfigurationGetRequest struct {
	Id string
}

// Information required to create a DeviceChain microservice.
type MicroserviceCreateRequest struct {
	Id              string
	InstanceId      string
	Name            string
	Description     string
	ConfigurationId string
}

// Information required to get a microservice.
type MicroserviceGetRequest struct {
	InstanceId     string
	MicroserviceId string
}

// Information required to list microservices.
type MicroserviceListRequest struct {
	InstanceId string
}

// -----------------------------
// Tenant Microservice Mangement
// -----------------------------

// Information required to create a DeviceChain tenant microservice.
type TenantMicroserviceCreateRequest struct {
	InstanceId     string
	TenantId       string
	MicroserviceId string
}

// Information required to get a tenant microservice.
type TenantMicroserviceGetRequest struct {
	InstanceId           string
	TenantMicroserviceId string
}

// Information required to get a tenant microservice.
type TenantMicroserviceByTenantRequest struct {
	InstanceId string
	TenantId   string
}
