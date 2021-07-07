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
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	k8sv1alpha1 "github.com/yozel/taint-operator/api/v1alpha1"
	"github.com/yozel/taint-operator/internal/taintset"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Taint controller", func() {
	const (
		TaintName = "test-taint"
		NodeName  = "test-node"

		timeout  = time.Second * 30
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	getNode := func(name string) *corev1.Node {
		ctx := context.Background()
		r := &corev1.Node{}
		Eventually(func() bool {
			err := k8sClient.Get(ctx, types.NamespacedName{Name: name, Namespace: ""}, r)
			if err != nil {
				return false
			}
			return true
		}, timeout, interval).Should(BeTrue())

		return r
	}

	createNode := func(name string, labels map[string]string) *corev1.Node {
		ctx := context.Background()
		Expect(k8sClient.Create(ctx, &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name:   name,
				Labels: labels,
			},
		})).Should(Succeed())
		return getNode(name)
	}

	createTaint := func(name string, nodeSelector map[string]string, taints []corev1.Taint) *k8sv1alpha1.Taint {
		ctx := context.Background()
		taintObject := &k8sv1alpha1.Taint{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: k8sv1alpha1.TaintSpec{
				NodeSelector: nodeSelector,
				Taints:       taints,
			},
		}
		Expect(k8sClient.Create(ctx, taintObject)).Should(Succeed())

		taintLookupKey := types.NamespacedName{Name: TaintName, Namespace: ""}
		r := &k8sv1alpha1.Taint{}

		// We'll need to retry getting this newly created CronJob, given that creation may not immediately happen.
		Eventually(func() bool {
			err := k8sClient.Get(ctx, taintLookupKey, r)
			if err != nil {
				return false
			}
			return true
		}, timeout, interval).Should(BeTrue())

		return r
	}

	Context("There is one Node", func() {
		BeforeEach(func() {
			createNode(NodeName, map[string]string{"component": "elasticsearch"})
		})

		It("should add a new taint to Node when a new Taint is created", func() {
			taints := []corev1.Taint{
				{
					Key:    "component",
					Value:  "elasticsearch",
					Effect: corev1.TaintEffectNoExecute,
				},
			}

			createTaint(TaintName, map[string]string{"component": "elasticsearch"}, taints)

			Eventually(func() *taintset.TaintSet {
				n := getNode(NodeName)
				return taintset.NewTaintSet(n.Spec.Taints)
			}, timeout, interval).Should(taintset.Contains(taintset.NewTaintSet(taints)))
		})
	})
})
