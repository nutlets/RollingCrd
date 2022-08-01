/*
Copyright 2022.
*/

package controllers

import (
	"context"
	"rollingCRD/pkg/monitor"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "rollingCRD/api/v1"
)

// RollingUpdateCrdReconciler reconciles a RollingUpdateCrd object
type RollingUpdateCrdReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const MonitorFinalizer = "transwarp.io.tos.monitored/finalizer"

//+kubebuilder:rbac:groups=demo.roll.io,resources=rollingupdatecrds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.roll.io,resources=rollingupdatecrds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.roll.io,resources=rollingupdatecrds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RollingUpdateCrd object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *RollingUpdateCrdReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	logger.Info("start reconcile")
	rollingUpdateCrd := &demov1.RollingUpdateCrd{}

	key := client.ObjectKey{Namespace: req.Namespace, Name: req.Name}
	if err := r.Get(ctx, key, rollingUpdateCrd); err != nil {
		logger.Error(err, "get target obj failed")
		return ctrl.Result{}, err
	}

	if rollingUpdateCrd.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(rollingUpdateCrd, MonitorFinalizer) {
			controllerutil.AddFinalizer(rollingUpdateCrd, MonitorFinalizer)
			if err := r.Update(ctx, rollingUpdateCrd); err != nil {
				logger.Error(err, "add MonitorFinalizer failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(rollingUpdateCrd, MonitorFinalizer) {
			monitor.RemoveMonitoredDeploy(rollingUpdateCrd.Namespace, rollingUpdateCrd.Spec.DeploymentName)
			controllerutil.RemoveFinalizer(rollingUpdateCrd, MonitorFinalizer)
			if err := r.Update(ctx, rollingUpdateCrd); err != nil {
				logger.Error(err, "remove MonitorFinalizer failed")
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	monitor.AddMonitorDeploy(rollingUpdateCrd.Namespace, rollingUpdateCrd.Spec.DeploymentName)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RollingUpdateCrdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.RollingUpdateCrd{}).
		Complete(r)
}
