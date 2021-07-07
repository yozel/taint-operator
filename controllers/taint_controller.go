/*
Copyright 2021.

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	k8sv1alpha1 "github.com/yozel/taint-operator/api/v1alpha1"
)

const taintFinalizer = "taint.k8s.yozel.co/finalizer"

// TaintReconciler reconciles a Taint object
type TaintReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=k8s.yozel.co,resources=taints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=k8s.yozel.co,resources=taints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=k8s.yozel.co,resources=taints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Taint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *TaintReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return reconcile(ctx, r, r.Log, req)
}

// SetupWithManager sets up the controller with the Manager.
func (r *TaintReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sv1alpha1.Taint{}).
		Complete(r)
}
