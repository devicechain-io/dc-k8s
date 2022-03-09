/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	ClientConfig  *rest.Config
	V1Beta1Client client.Client
	V1Client      client.Client
)

func initClientConfig() {
	ClientConfig = config.GetConfigOrDie()
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

func initV1Client() error {
	scheme := runtime.NewScheme()
	err := v1.SchemeBuilder.AddToScheme(scheme)
	if err != nil {
		return err
	}

	V1Client, err = client.New(ClientConfig, client.Options{Scheme: scheme})
	if err != nil {
		return err
	}
	return nil
}

// Create a new namespace
func createNamespace(nsid string) (*v1.Namespace, error) {
	ns := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: nsid}}

	// Attempt to create the namespace.
	err := V1Client.Create(context.Background(), ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// Get namespace by id
func getNamespace(nsid string) (*v1.Namespace, error) {
	ns := &v1.Namespace{}
	err := V1Client.Get(context.Background(), client.ObjectKey{
		Name: nsid,
	}, ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
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
	config, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	instance := &Instance{
		ObjectMeta: metav1.ObjectMeta{
			Name: request.Id,
		},
		Spec: InstanceSpec{
			Name:          request.Name,
			Description:   request.Description,
			Configuration: EntityConfiguration{RawMessage: config},
		},
	}

	// Attempt to create the instance.
	err = V1Beta1Client.Create(context.Background(), instance)
	if err != nil {
		return nil, err
	}

	// Locate namespace same as instance id and create if not existing
	_, err = getNamespace(request.Id)
	if err != nil {
		_, err = createNamespace(request.Id)
		if err != nil {
			return nil, err
		}
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

// Information required to create a DeviceChain tenant.
type TenantCreateRequest struct {
	InstanceId  string
	Id          string
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
			Name:      request.Id,
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
		Name:      request.Id,
		Namespace: instance.GetObjectMeta().GetName(),
	}, tenant)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func init() {
	initClientConfig()
	err := initV1Beta1Client()
	if err != nil {
		log.Fatal("unable to initialize v1beta1 client", err)
	}
	err = initV1Client()
	if err != nil {
		log.Fatal("unable to initialize v1 client", err)
	}
}