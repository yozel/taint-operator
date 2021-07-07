package taintset

import (
	"crypto"
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func Hash(t corev1.Taint) string {
	digester := crypto.SHA256.New()
	fmt.Fprintf(digester, "%s\n%s\n%s\n", t.Key, t.Value, t.Effect)
	return string(digester.Sum(nil))
}

type TaintSet struct {
	items map[string]corev1.Taint
}

func (t *TaintSet) Copy() *TaintSet {
	c := &TaintSet{items: make(map[string]corev1.Taint, len(t.items))}
	for k, taint := range t.items {
		c.items[k] = taint
	}
	return c
}

func (t *TaintSet) Len() int {
	return len(t.items)
}

func (t *TaintSet) Add(items ...corev1.Taint) {
	for _, taint := range items {
		t.items[Hash(taint)] = taint
	}
}

func (t *TaintSet) Has(item corev1.Taint) bool {
	_, ok := t.items[Hash(item)]
	return ok
}

func (t *TaintSet) IsSubsetOf(t2 *TaintSet) bool {
	for k, _ := range t.items {
		_, ok := t2.items[k]
		if !ok {
			return false
		}
	}
	return true
}

func (t *TaintSet) IsSupersetOf(t2 *TaintSet) bool {
	return t2.IsSubsetOf(t)
}

func NewTaintSet(items []corev1.Taint) *TaintSet {
	t := &TaintSet{items: make(map[string]corev1.Taint, len(items))}
	for _, taint := range items {
		t.items[Hash(taint)] = taint
	}
	return t
}
