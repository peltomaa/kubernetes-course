/*
Copyright 2025.

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
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	dummysitev1 "dummy-site/api/v1"
)

// DummySiteReconciler reconciles a DummySite object
type DummySiteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=dummy-site.dummy-site,resources=dummysites,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dummy-site.dummy-site,resources=dummysites/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dummy-site.dummy-site,resources=dummysites/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DummySite object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.0/pkg/reconcile
func (r *DummySiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

func HTTPServer() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Log.Error(err, "Failed to start server on :8080")
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *DummySiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	go HTTPServer()

	return ctrl.NewControllerManagedBy(mgr).
		For(&dummysitev1.DummySite{}).
		Named("dummysite").
		Complete(r)
}
