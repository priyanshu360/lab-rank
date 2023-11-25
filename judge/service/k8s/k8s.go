package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// KubernetesManager is a struct that encapsulates the logic for interacting with Kubernetes.
type KubernetesManager struct {
	Clientset *kubernetes.Clientset
	Namespace string
}

// NewKubernetesManager creates a new instance of KubernetesManager.
func NewKubernetesManager(kubeconfig string, namespace string) (*KubernetesManager, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating clientset: %v", err)
	}

	return &KubernetesManager{
		Clientset: clientset,
		Namespace: namespace,
	}, nil
}

// Init initializes any necessary resources in Kubernetes.
func (k8s *KubernetesManager) Init() error {
	// Add any initialization logic here.
	// For example, creating ConfigMaps, Secrets, etc.

	return nil
}
