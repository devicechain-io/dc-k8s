/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package controllers

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/devicechain-io/dc-k8s/api/v1beta1"
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenants/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenants/finalizers,verbs=update
func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	tenant := &v1beta1.Tenant{}
	err := r.Get(ctx, req.NamespacedName, tenant)
	if err != nil {
		return ctrl.Result{}, err
	}

	log.Printf("Tenant reconcile for ctx: %v request: %v", ctx, req)
	if tenant != nil {
		tmslist, err := v1beta1.GetTenantMicroservicesForTenant(v1beta1.TenantMicroserviceByTenantRequest{
			InstanceId: tenant.GetObjectMeta().GetNamespace(),
			TenantId:   tenant.GetObjectMeta().GetName()})
		if err != nil {
			return ctrl.Result{}, err
		}

		tmsbymsid := map[string]v1beta1.TenantMicroservice{}
		for _, tms := range tmslist.Items {
			tmsbymsid[tms.Spec.MicroserviceId] = tms
		}
		log.Printf("Found tms by id: %v\n", tmsbymsid)

		mslist, err := v1beta1.ListMicroservices(v1beta1.MicroserviceListRequest{
			InstanceId: tenant.GetObjectMeta().GetNamespace(),
		})
		if err != nil {
			return ctrl.Result{}, err
		}
		for _, ms := range mslist.Items {
			if _, present := tmsbymsid[ms.ObjectMeta.Name]; !present {
				tms, err := addTenantMicroservice(tenant, ms)
				if err != nil {
					return ctrl.Result{}, err
				}
				log.Printf("Created tenant microservice: %v", tms)
			}
		}
	}

	return ctrl.Result{}, nil
}

// Creates a tenant microservice for a tenant/microservice combination.
func addTenantMicroservice(tenant *v1beta1.Tenant, ms v1beta1.Microservice) (*v1beta1.TenantMicroservice, error) {
	return v1beta1.CreateTenantMicroservice(v1beta1.TenantMicroserviceCreateRequest{
		InstanceId:     tenant.GetObjectMeta().GetNamespace(),
		TenantId:       tenant.ObjectMeta.Name,
		MicroserviceId: ms.ObjectMeta.Name})
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.Tenant{}).
		Complete(r)
}
