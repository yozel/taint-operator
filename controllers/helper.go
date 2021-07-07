package controllers

import (
	"context"
	"sync"

	"github.com/go-logr/logr"
	k8sv1alpha1 "github.com/yozel/taint-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type genericReconciler interface {
	client.Reader
	client.Writer
}

var mu sync.Mutex

func reconcile(ctx context.Context, r genericReconciler, log logr.Logger, req ctrl.Request) (res ctrl.Result, err error) {
	mu.Lock()
	log.Info("Reconciliation started")
	nodes := &corev1.NodeList{}
	err = r.List(ctx, nodes)
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}
	taints := &k8sv1alpha1.TaintList{}
	err = r.List(ctx, taints)
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}
	for _, node := range nodes.Items {
		log := log.WithValues("node", node.GetName())
		node.Spec.Taints = []corev1.Taint{}
	TaintLoop:
		for _, taint := range taints.Items {
			log := log.WithValues("taint", taint.GetName())
			for labelKey, labelValue := range taint.Spec.NodeSelector {
				if node.Labels[labelKey] != labelValue {
					continue TaintLoop
				}
			}
			log.Info("Matched a taint and a node")
			node.Spec.Taints = append(node.Spec.Taints, taint.Spec.DeepCopy().Taints...)
		}
		log.WithValues("node taints", node.Spec.Taints).Info("Updating Node with taints")
		r.Update(ctx, &node)
	}
	log.Info("Reconciliation done")
	mu.Unlock()
	return ctrl.Result{}, nil
}
