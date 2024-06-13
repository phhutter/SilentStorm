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
	"time"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	silentstormv1alpha1 "github.com/biggold1310/silentstorm/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/equality"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SilenceReconciler reconciles a Alertmanager object
type SilenceReconciler struct {
	SharedReconciler
}

//+kubebuilder:rbac:groups=silentstorm.biggold1310.ch,resources=silences,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=silentstorm.biggold1310.ch,resources=silences/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=silentstorm.biggold1310.ch,resources=silences/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SilenceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var err error
	_ = log.FromContext(ctx)

	silence := &silentstormv1alpha1.Silence{}
	err = r.Client.Get(ctx, req.NamespacedName, silence)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !controllerutil.ContainsFinalizer(silence, finalizer) {
		controllerutil.AddFinalizer(silence, finalizer)
		return ctrl.Result{}, r.Client.Update(ctx, silence)
	}

	oldStatus := silence.Status.DeepCopy()
	oldMetadata := silence.ObjectMeta.DeepCopy()

	// Add matcher for the actual namespace
	alertmanagerSilence := silence.Spec.AlertmanagerSilence
	exists := false
	for _, matcher := range alertmanagerSilence.Matchers {
		if matcher.Name == "namespace" {
			exists = true
			matcher.IsRegex = false
			matcher.IsEqual = true
			matcher.Value = silence.GetNamespace()
		}
	}
	if !exists {
		alertmanagerSilence.Matchers = append(alertmanagerSilence.Matchers, silentstormv1alpha1.Matcher{
			IsEqual: true,
			IsRegex: false,
			Name:    "namespace",
			Value:   silence.GetNamespace(),
		})
	}

	err = r.handleSilence(ctx, &silence.ObjectMeta, alertmanagerSilence, &silence.Status.AlertmanagerReferences)
	if err != nil {
		return ctrl.Result{}, err
	}

	// We check the deletion status after the processing as the processing is deleting based on the GetDeletionTimestamp
	if silence.GetDeletionTimestamp() != nil && isClean(&silence.Status.AlertmanagerReferences) {
		controllerutil.RemoveFinalizer(silence, finalizer)
		return ctrl.Result{}, r.Client.Update(ctx, silence)
	}

	if !equality.Semantic.DeepEqual(oldMetadata, &silence.ObjectMeta) {
		// Update whole object if MetaData changes (this should only happen for deletions)
		return ctrl.Result{}, r.Client.Update(ctx, silence)
	} else if !equality.Semantic.DeepEqual(oldStatus, &silence.Status) {
		// Regular update of the Status section of the object
		return ctrl.Result{}, r.Client.Status().Update(ctx, silence)
	}

	return ctrl.Result{RequeueAfter: 1 * time.Hour}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SilenceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&silentstormv1alpha1.Silence{}).
		Complete(r)
}
