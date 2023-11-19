package k8s

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

// CreateJob creates a Kubernetes Job.
func (k8s *KubernetesManager) CreateJob(job *batchv1.Job) (*batchv1.Job, error) {
	resultJob, err := k8s.Clientset.BatchV1().Jobs(k8s.Namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error creating Job: %v", err)
	}

	return resultJob, nil
}

func main() {
	kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "path to kubeconfig file")
	namespace := flag.String("namespace", "default", "Kubernetes namespace")
	flag.Parse()

	k8s, err := NewKubernetesManager(*kubeconfig, *namespace)
	if err != nil {
		fmt.Printf("Error creating KubernetesManager: %v\n", err)
		os.Exit(1)
	}

	// Initialize any necessary resources
	err = k8s.Init()
	if err != nil {
		fmt.Printf("Error initializing Kubernetes resources: %v\n", err)
		os.Exit(1)
	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-job",
			Namespace: *namespace,
		},
		Spec: batchv1.JobSpec{
			Template: batchv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "example",
					},
				},
				Spec: batchv1.PodSpec{
					RestartPolicy: "Never",
					Containers: []corev1.Container{
						{
							Name:  "example-container",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}

	// Create the Job
	resultJob, err := k8s.CreateJob(job)
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Job %q created\n", resultJob.GetObjectMeta().GetName())
}
