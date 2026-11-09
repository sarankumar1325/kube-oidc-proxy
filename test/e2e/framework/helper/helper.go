// Copyright Jetstack Ltd. See LICENSE for details.
package helper

import (
	"k8s.io/client-go/kubernetes"

	"github.com/rafpe/kube-oidc-proxy/test/e2e/framework/config"
)

// Helper provides methods for common operations needed during tests.
type Helper struct {
	cfg *config.Config

	// KubeClient talks to the test cluster directly, as its admin, rather
	// than through the proxy.
	KubeClient kubernetes.Interface
}

// NewHelper returns a Helper for the cluster cfg describes, which kubeClient
// must already be connected to.
func NewHelper(cfg *config.Config, kubeClient kubernetes.Interface) *Helper {
	return &Helper{
		cfg:        cfg,
		KubeClient: kubeClient,
	}
}

// RepoRoot is the repository checkout the suite runs from, for helpers that
// render the chart.
func (h *Helper) RepoRoot() string {
	return h.cfg.RepoRoot
}
