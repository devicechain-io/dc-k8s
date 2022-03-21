/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package controllers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/devicechain-io/dc-k8s/api/v1beta1"
)

const (
	ENV_TENANT_ID          = "DC_TENANT_ID"
	ENV_TENANT_NAME        = "DC_TENANT_NAME"
	ENV_MICROSERVICE_ID    = "DC_MICROSERVICE_ID"
	ENV_MICROSERVICE_NAME  = "DC_MICROSERVICE_NAME"
	ENV_MS_FUNCTIONAL_AREA = "DC_MS_FUNCTIONAL_AREA"
)

// TenantMicroserviceReconciler reconciles a TenantMicroservice object
type TenantMicroserviceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenantmicroservices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenantmicroservices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=core.devicechain.io,resources=tenantmicroservices/finalizers,verbs=update
func (r *TenantMicroserviceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	tms := &v1beta1.TenantMicroservice{}
	err := r.Get(ctx, req.NamespacedName, tms)
	if err != nil {
		log.Info(fmt.Sprintf("Handling deleted tenant microservice: %+v", req.NamespacedName))
		err := r.handleTenantMicroserviceDeleted(ctx, req)
		if err != nil {
			log.Error(err, "Unable to handle tenant microservice delete")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// Handle creating or updating a k8s Deployment for the tenant microservice
	log.Info(fmt.Sprintf("Handling added/updated tenant microservice: %+v", req.NamespacedName))
	err = r.createOrUpdateDeployment(ctx, tms)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Update tenant config map entry with configuration
	err = r.updateTenantConfigMap(ctx, tms)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantMicroserviceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.TenantMicroservice{}).
		Complete(r)
}

// Get namespaced name for deployment
func getDeploymentName(tms *v1beta1.TenantMicroservice) types.NamespacedName {
	return types.NamespacedName{Namespace: tms.ObjectMeta.Namespace, Name: tms.ObjectMeta.Name}
}

// Create or update a k8s Deployment for the tenant microservice
func (r *TenantMicroserviceReconciler) createOrUpdateDeployment(ctx context.Context, tms *v1beta1.TenantMicroservice) error {
	log := logf.FromContext(ctx)

	// Attempt to look up existing deployment.
	dname := getDeploymentName(tms)
	deploy := &appsv1.Deployment{}
	err := r.Get(ctx, dname, deploy)
	if err != nil {
		log.Info(fmt.Sprintf("Existing deployment not found for tenant microservice: %+v", dname))

		// Look up associated microservice.
		dct, err := v1beta1.GetTenant(v1beta1.TenantGetRequest{
			InstanceId: tms.ObjectMeta.Namespace,
			TenantId:   tms.Spec.TenantId,
		})
		if err != nil {
			return err
		}

		// Look up associated microservice.
		ms, err := v1beta1.GetMicroservice(v1beta1.MicroserviceGetRequest{
			InstanceId:     tms.ObjectMeta.Namespace,
			MicroserviceId: tms.Spec.MicroserviceId,
		})
		if err != nil {
			return err
		}

		// Create a new deployment.
		_, err = r.createDeployment(ctx, tms, dct, ms)
		return err
	}

	log.Info(fmt.Sprintf("Existing deployment found for tenant microservice: %+v", dname))

	return nil
}

// Create labels to target deployment
func createDeploymentLabels(tms *v1beta1.TenantMicroservice) map[string]string {
	return map[string]string{
		v1beta1.LABEL_TENANT:       tms.Spec.TenantId,
		v1beta1.LABEL_MICROSERVICE: tms.Spec.MicroserviceId,
	}
}

// Create a deployment based on tenant microservice details
func (r *TenantMicroserviceReconciler) createDeployment(ctx context.Context, tms *v1beta1.TenantMicroservice,
	dct *v1beta1.Tenant, ms *v1beta1.Microservice) (*appsv1.Deployment, error) {
	log := logf.FromContext(ctx)

	dci, err := v1beta1.GetInstance(v1beta1.InstanceGetRequest{Id: tms.ObjectMeta.Namespace})
	if err != nil {
		return nil, err
	}

	dname := getDeploymentName(tms)
	labels := createDeploymentLabels(tms)
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dname.Name,
			Namespace: dname.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            tms.Spec.MicroserviceId,
							Image:           ms.Spec.Image,
							ImagePullPolicy: ms.Spec.ImagePullPolicy,
							Env: []corev1.EnvVar{
								{
									Name:  ENV_TENANT_ID,
									Value: dct.ObjectMeta.Name,
								},
								{
									Name:  ENV_TENANT_NAME,
									Value: dct.Spec.Name,
								},
								{
									Name:  ENV_MICROSERVICE_ID,
									Value: ms.ObjectMeta.Name,
								},
								{
									Name:  ENV_MICROSERVICE_NAME,
									Value: ms.Spec.Name,
								},
								{
									Name:  ENV_MS_FUNCTIONAL_AREA,
									Value: ms.Spec.FunctionalArea,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "instance-config",
									MountPath: "/etc/dci-config",
								}, {
									Name:      "tenant-config",
									MountPath: "/etc/dct-config",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "instance-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: getInstanceConfigMapName(dci),
									},
								},
							},
						},
						{
							Name: "tenant-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: getTenantConfigMapName(tms.Spec.TenantId),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	err = r.Create(context.Background(), deploy)
	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("Created k8s Deployment for tenant microservice: %+v", dname))
	return deploy, nil
}

// Handle a deleted tenant microservice
func (r *TenantMicroserviceReconciler) handleTenantMicroserviceDeleted(ctx context.Context, req ctrl.Request) error {
	log := logf.FromContext(ctx)

	deploy := &appsv1.Deployment{}
	err := r.Get(ctx, req.NamespacedName, deploy)
	if err != nil {
		log.Info(fmt.Sprintf("Unable to find deployment for tenant microservice: %+v", req.NamespacedName))
		return err
	}
	err = r.Delete(ctx, deploy)
	log.Info(fmt.Sprintf("Deleted deployment for tenant microservice: %+v", req.NamespacedName))
	return err
}

// Update tenant configuration map with entry for tenant microservice
func (r *TenantMicroserviceReconciler) updateTenantConfigMap(ctx context.Context,
	tms *v1beta1.TenantMicroservice) error {

	// Get microservice information.
	ms, err := v1beta1.GetMicroservice(v1beta1.MicroserviceGetRequest{
		InstanceId:     tms.ObjectMeta.Namespace,
		MicroserviceId: tms.Spec.MicroserviceId,
	})
	if err != nil {
		return err
	}

	tcmap, err := getTenantConfigMap(tms.Spec.TenantId, tms.ObjectMeta.Namespace)
	if err != nil {
		return err
	}

	// Make sure map data is populated.
	if tcmap.Data == nil {
		tcmap.Data = make(map[string]string, 0)
	}

	// Update map index for functional area
	tcmap.Data[ms.Spec.FunctionalArea] = string(tms.Spec.Configuration.RawMessage)
	err = r.Update(ctx, tcmap)
	return err
}
