/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package v1beta1

import (
	"context"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func NewClient() (client.Client, error) {
	v1beta1, err := SchemeBuilder.Build()
	if err != nil {
		return nil, err
	}
	client, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: v1beta1})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Information required to create a DeviceChain instance.
type InstanceCreateRequest struct {
	Id              string
	Name            string
	Description     string
	ConfigurationId string
}

// Create a new DeviceChain instance CR.
func CreateInstance(cli client.Client, request InstanceCreateRequest) (*Instance, error) {
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
	err = cli.Create(context.Background(), instance)
	if err != nil {
		return nil, err
	}

	// Attempt to get the created instance.
	err = cli.Get(context.Background(), client.ObjectKey{
		Name: request.Id,
	}, instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
