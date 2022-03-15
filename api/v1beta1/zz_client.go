/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

import (
	"context"
	"errors"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	LABEL_TENANT       = "devicechain.io.tenant"
	LABEL_MICROSERVICE = "devicechain.io.microservice"
)

var (
	ClientConfig  *rest.Config
	V1Beta1Client client.Client
)

func initClientConfig() {
	ClientConfig = config.GetConfigOrDie()
	ClientConfig.RateLimiter = flowcontrol.NewFakeAlwaysRateLimiter()
}

func initV1Beta1Client() error {
	v1beta1, err := SchemeBuilder.Build()
	if err != nil {
		return err
	}
	V1Beta1Client, err = client.New(ClientConfig, client.Options{Scheme: v1beta1})
	if err != nil {
		return err
	}
	return nil
}

// Get instance configuraion by id
func getInstanceConfiguration(id string) (*InstanceConfiguration, error) {
	ic := &InstanceConfiguration{}
	err := V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name: id,
	}, ic)
	if err != nil {
		return nil, err
	}
	return ic, nil
}

// Information required to create a DeviceChain instance.
type InstanceCreateRequest struct {
	Id              string
	Name            string
	Description     string
	ConfigurationId string
}

// Create a new DeviceChain instance CR.
func CreateInstance(request InstanceCreateRequest) (*Instance, error) {
	ic, err := getInstanceConfiguration(request.ConfigurationId)
	if err != nil {
		return nil, err
	}

	instance := &Instance{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.Id,
		},
		Spec: InstanceSpec{
			Name:            request.Name,
			Description:     request.Description,
			ConfigurationId: request.ConfigurationId,
			Configuration:   EntityConfiguration{RawMessage: ic.Spec.Configuration.RawMessage},
		},
	}

	// Attempt to create the instance.
	err = V1Beta1Client.Create(context.Background(), instance)
	if err != nil {
		return nil, err
	}

	// Attempt to get the created instance.
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name: request.Id,
	}, instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Information required to get a DeviceChain instance.
type InstanceGetRequest struct {
	Id string
}

// Get an instance based on request criteria
func GetInstance(request InstanceGetRequest) (*Instance, error) {
	instance := &Instance{}
	err := V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name: request.Id,
	}, instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Information required to get a tenant.
type TenantGetRequest struct {
	InstanceId string
	TenantId   string
}

// Get a tenant based on request criteria
func GetTenant(request TenantGetRequest) (*Tenant, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	tenant := &Tenant{}
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      request.TenantId,
		Namespace: instance.GetObjectMeta().GetName(),
	}, tenant)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

// Information required to create a DeviceChain tenant.
type TenantCreateRequest struct {
	InstanceId  string
	TenantId    string
	Name        string
	Description string
}

// Create a new DeviceChain tenant CR.
func CreateTenant(request TenantCreateRequest) (*Tenant, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided when creating tenant")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	// Create tenant in instance namespace
	tenant := &Tenant{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.TenantId,
			Namespace: instance.GetObjectMeta().GetName(),
		},
		Spec: TenantSpec{
			Name:        request.Name,
			Description: request.Description,
		},
	}

	// Attempt to create the tenant.
	err = V1Beta1Client.Create(context.Background(), tenant)
	if err != nil {
		return nil, err
	}

	// Attempt to get the created tenant.
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      request.TenantId,
		Namespace: instance.GetObjectMeta().GetName(),
	}, tenant)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

// Information required to get a microservice configuration.
type MicroserviceConfigurationGetRequest struct {
	Id string
}

// Get an microservice configuration based on request criteria
func GetMicroserviceConfiguration(request MicroserviceConfigurationGetRequest) (*MicroserviceConfiguration, error) {
	msconfig := &MicroserviceConfiguration{}
	err := V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name: request.Id,
	}, msconfig)
	if err != nil {
		return nil, err
	}
	return msconfig, nil
}

// Information required to get a microservice.
type MicroservicGetRequest struct {
	InstanceId     string
	MicroserviceId string
}

// Get a microservice based on request criteria
func GetMicroservice(request MicroservicGetRequest) (*Microservice, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	ms := &Microservice{}
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      request.MicroserviceId,
		Namespace: instance.GetObjectMeta().GetName(),
	}, ms)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

// Information required to create a DeviceChain microservice.
type MicroserviceCreateRequest struct {
	Id              string
	InstanceId      string
	Name            string
	Description     string
	ConfigurationId string
}

// Create a new DeviceChain microservice CR.
func CreateMicroservice(request MicroserviceCreateRequest) (*Microservice, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided when creating microservice")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	if request.ConfigurationId == "" {
		return nil, errors.New("configuration id must be provided when creating microservice")
	}
	msc, err := GetMicroserviceConfiguration(MicroserviceConfigurationGetRequest{Id: request.ConfigurationId})
	if err != nil {
		return nil, err
	}

	// Create ms in instance namespace
	ms := &Microservice{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Id,
			Namespace: instance.GetObjectMeta().GetName(),
		},
		Spec: MicroserviceSpec{
			Name:            request.Name,
			Description:     request.Description,
			FunctionalArea:  msc.Spec.FunctionalArea,
			Image:           msc.Spec.Image,
			ImagePullPolicy: v1.PullAlways,
			ConfigurationId: request.ConfigurationId,
			Configuration:   EntityConfiguration{RawMessage: msc.Spec.Configuration.RawMessage},
		},
	}

	// Attempt to create the microservice.
	err = V1Beta1Client.Create(context.Background(), ms)
	if err != nil {
		return nil, err
	}

	// Attempt to get the created microservice.
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      request.Id,
		Namespace: instance.GetObjectMeta().GetName(),
	}, ms)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

// Information required to list microservices.
type MicroserviceListRequest struct {
	InstanceId string
}

// List microservices that match the given criteria
func ListMicroservices(request MicroserviceListRequest) (*MicroserviceList, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	mslist := &MicroserviceList{}
	err = V1Beta1Client.List(context.Background(), mslist, client.InNamespace(instance.GetObjectMeta().GetName()))
	if err != nil {
		return nil, err
	}
	return mslist, nil
}

// Information required to get a tenant microservice.
type TenantMicroserviceGetRequest struct {
	InstanceId     string
	MicroserviceId string
}

// Get a tenant microservice based on request criteria
func GetTenantMicroservice(request TenantMicroserviceGetRequest) (*TenantMicroservice, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	tms := &TenantMicroservice{}
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      request.MicroserviceId,
		Namespace: instance.GetObjectMeta().GetName(),
	}, tms)
	if err != nil {
		return nil, err
	}
	return tms, nil
}

// Information required to get a tenant microservice.
type TenantMicroserviceByTenantRequest struct {
	InstanceId string
	TenantId   string
}

// Get a tenant microservice based on request criteria
func GetTenantMicroservicesForTenant(request TenantMicroserviceByTenantRequest) (*TenantMicroserviceList, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	// Verify tenant exists
	_, err = GetTenant(TenantGetRequest{
		InstanceId: instance.GetObjectMeta().GetName(),
		TenantId:   request.TenantId})
	if err != nil {
		return nil, err
	}

	// List tenant microservices in instance namespace with tenant label
	tmslist := &TenantMicroserviceList{}
	err = V1Beta1Client.List(context.Background(), tmslist, client.InNamespace(instance.GetObjectMeta().GetName()),
		client.MatchingLabels{LABEL_TENANT: request.TenantId})
	if err != nil {
		return nil, err
	}
	return tmslist, nil
}

// Information required to create a DeviceChain tenant microservice.
type TenantMicroserviceCreateRequest struct {
	InstanceId     string
	TenantId       string
	MicroserviceId string
}

// Create a new tenant microservice CR.
func CreateTenantMicroservice(request TenantMicroserviceCreateRequest) (*TenantMicroservice, error) {
	if request.InstanceId == "" {
		return nil, errors.New("instance id must be provided")
	}
	instance, err := GetInstance(InstanceGetRequest{Id: request.InstanceId})
	if err != nil {
		return nil, err
	}

	if request.TenantId == "" {
		return nil, errors.New("tenant id must be provided when creating tenant microservice")
	}
	tenant, err := GetTenant(TenantGetRequest{
		InstanceId: instance.GetObjectMeta().GetName(),
		TenantId:   request.TenantId})
	if err != nil {
		return nil, err
	}

	if request.MicroserviceId == "" {
		return nil, errors.New("microservice id must be provided when creating tenant microservice")
	}
	ms, err := GetMicroservice(MicroservicGetRequest{
		InstanceId:     instance.GetObjectMeta().GetName(),
		MicroserviceId: request.MicroserviceId})
	if err != nil {
		return nil, err
	}

	// Create tenant ms in instance namespace
	tmsid := fmt.Sprintf("%s-%s-%s", "tms", tenant.ObjectMeta.Name, ms.ObjectMeta.Name)
	tms := &TenantMicroservice{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tmsid,
			Namespace: tenant.GetObjectMeta().GetNamespace(),
			Labels: map[string]string{
				LABEL_TENANT:       tenant.GetObjectMeta().GetName(),
				LABEL_MICROSERVICE: ms.GetObjectMeta().GetName(),
			},
		},
		Spec: TenantMicroserviceSpec{
			MicroserviceId: request.MicroserviceId,
			TenantId:       request.TenantId,
			Configuration:  EntityConfiguration{RawMessage: ms.Spec.Configuration.RawMessage},
		},
	}

	// Attempt to create the tenant microservice.
	err = V1Beta1Client.Create(context.Background(), tms)
	if err != nil {
		return nil, err
	}

	// Attempt to get the created microservice.
	err = V1Beta1Client.Get(context.Background(), client.ObjectKey{
		Name:      tmsid,
		Namespace: tenant.GetObjectMeta().GetNamespace(),
	}, tms)
	if err != nil {
		return nil, err
	}
	return tms, nil
}

func init() {
	initClientConfig()
	err := initV1Beta1Client()
	if err != nil {
		log.Fatal("unable to initialize v1beta1 client", err)
	}
}
