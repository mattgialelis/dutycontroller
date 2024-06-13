/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	pagerdutyv1beta1 "github.com/mattgialelis/dutycontroller/api/v1beta1"
	"github.com/mattgialelis/dutycontroller/pkg/providers/pd"
)

const (
	serviceFinalizer = "service.dutycontroller.io/finalizer"
)

// ServicesReconciler reconciles a Services object
type ServicesReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	PagerClient *pd.Pagerduty
	recorder    record.EventRecorder
}

//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=services/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=services/finalizers,verbs=update

func (r *ServicesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.WithValues("service", req.NamespacedName)
	var service pagerdutyv1beta1.Services
	var businessService pagerdutyv1beta1.BusinessService
	var businessSeriviceId string

	// Fetch the Service instance
	if err := r.Get(ctx, req.NamespacedName, &service); err != nil {
		//do not requeue if the resource does not exist
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get resource: %w", err)
	}

	//Defers the update of the Service status
	defer func() {
		if service.DeletionTimestamp.IsZero() {
			if err := r.Status().Update(ctx, &service); err != nil {
				log.Error(err, "unable to update Service status")
			}
		}
	}()

	//Lookups the BusinessService to get its ID first looks in the cluster and if not found fetches from PagerDuty
	if service.Spec.BusinessService != "" {
		if err := r.Get(ctx, client.ObjectKey{
			Namespace: req.Namespace,
			Name:      service.Spec.BusinessService,
		}, &businessService); err != nil {
			if apierrors.IsNotFound(err) {
				log.Info("BusinessService not found in cluster, fetching from PagerDuty", "namespace", req.Namespace, "name", service.Spec.BusinessService)
				businessSeriviceId, _, err = r.PagerClient.GetBusinessServicebyName(req.Name)
				if err != nil {
					return ctrl.Result{}, fmt.Errorf("could not get business service by name: %w", err)
				}
			} else {
				return ctrl.Result{}, fmt.Errorf("get BusinessService: %w", err)
			}
		} else {
			businessSeriviceId = businessService.Status.ID
		}
	}

	//Check if the Service instance exists
	_, exists, err := r.PagerClient.GetPagerDutyServiceByNameDirect(req.Name)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("could not get service by name: %w", err)
	}

	//Get escalation policy ID
	escalationPolicyId, _, err := r.PagerClient.GetEscalationPolicyByName(service.Spec.EscalationPolicy)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("could not get escalation policy by name: %w", err)
	}

	//Convert the Service CRD to a Service struct
	pagerService := pd.ServicesSpectoService(service, escalationPolicyId)

	if !exists {
		id, err := r.PagerClient.CreatePagerDutyService(pagerService)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("could not create service: %w", err)
		}

		//Wait for the service to be created
		time.Sleep(1 * time.Second)
		if businessSeriviceId != "" {
			err := r.PagerClient.AssociateServiceBusiness(id, businessSeriviceId)
			if err != nil {
				log.Error(err, "could not associate service with business service", "businessSeriviceId", businessSeriviceId, "serviceId", service.Status.ID)
			}
		}

		service.Status.ID = id
		return ctrl.Result{}, nil
	}

	//Update the Service instance
	err = r.PagerClient.UpdatePagerDutyService(pagerService)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("could not update service: %w", err)
	}

	if businessSeriviceId != "" {
		err := r.PagerClient.AssociateServiceBusiness(service.Status.ID, businessSeriviceId)
		if err != nil {
			log.Error(err, "could not associate service with business service", "businessSeriviceId", businessSeriviceId, "serviceId", service.Status.ID)
		}
	}

	// Check if the BusinessService instance is marked for deletion
	if service.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(&service, serviceFinalizer) {

			controllerutil.AddFinalizer(&service, serviceFinalizer)
			if err := r.Update(ctx, &service); err != nil {
				return ctrl.Result{}, fmt.Errorf("could not update finalizers: %w", err)
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(&service, serviceFinalizer) {

			log.Info("Deleting Service", "ID", service.Status.ID, "Name", service.Name)
			err := r.PagerClient.DeletePagerDutyService(service.Status.ID)
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("could not delete pagerduty service: %w", err)
			}

			// Remove the finalizer
			controllerutil.RemoveFinalizer(&service, serviceFinalizer)
			if err := r.Client.Update(ctx, &service, &client.UpdateOptions{}); err != nil {
				return ctrl.Result{}, fmt.Errorf("could not update finalizers: %w", err)
			}

			return ctrl.Result{}, nil
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServicesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("services-controller")

	return ctrl.NewControllerManagedBy(mgr).
		For(&pagerdutyv1beta1.Services{}).
		Complete(r)
}
