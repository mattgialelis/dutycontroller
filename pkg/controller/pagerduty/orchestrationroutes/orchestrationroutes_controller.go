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
	orchestrationRouteFinalizer = "orchestrationRoute.dutycontroller.io/finalizer"
)

// OrchestrationroutesReconciler reconciles a Orchestrationroutes object
type OrchestrationroutesReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	PagerClient *pd.Pagerduty
	recorder    record.EventRecorder
}

//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=orchestrationroutes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=orchestrationroutes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pagerduty.dutycontroller.io,resources=orchestrationroutes/finalizers,verbs=update

func (r *OrchestrationroutesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.WithValues("orchestrationRoutes", req.NamespacedName)

	var orchestrationRoute pagerdutyv1beta1.Orchestrationroutes

	// Fetch the BusinessService instance
	if err := r.Get(ctx, req.NamespacedName, &orchestrationRoute); err != nil {
		//do not requeue if the resource does not exist
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, fmt.Errorf("get resource: %w", err)
	}

	defer func() {
		if orchestrationRoute.DeletionTimestamp.IsZero() {
			if err := r.Status().Update(ctx, &orchestrationRoute); err != nil {
				log.Error(err, "unable to update orchestration Route status")
			}
		}
	}()

	// Check if the BusinessService instance is marked for deletion
	if orchestrationRoute.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(&orchestrationRoute, orchestrationRouteFinalizer) {

			controllerutil.AddFinalizer(&orchestrationRoute, orchestrationRouteFinalizer)
			if err := r.Update(ctx, &orchestrationRoute); err != nil {
				return ctrl.Result{}, fmt.Errorf("could not update finalizers: %w", err)
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(&orchestrationRoute, orchestrationRouteFinalizer) {
			// Run finalization logic for BusinessService
			// If the finalization logic fails, don't remove the finalizer so
			// that we can retry during the next reconciliation.

			err := r.PagerClient.DeleteBusinessService(orchestrationRoute.Status.ID)
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("could not delete orchestration Route: %w", err)
			}

			// Remove the finalizer
			controllerutil.RemoveFinalizer(&orchestrationRoute, orchestrationRouteFinalizer)
			if err := r.Client.Update(ctx, &orchestrationRoute, &client.UpdateOptions{}); err != nil {
				return ctrl.Result{}, fmt.Errorf("could not update finalizers: %w", err)
			}

			return ctrl.Result{}, nil
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OrchestrationroutesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("orchestrationroutes-controller")
	return ctrl.NewControllerManagedBy(mgr).
		For(&pagerdutyv1beta1.Orchestrationroutes{}).
		Complete(r)
}
