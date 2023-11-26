package executer

import (
	"context"
	"fmt"
	"log"

	"github.com/priyanshu360/lab-rank/judge/service/k8s"
	"github.com/priyanshu360/lab-rank/judge/utils"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Executer struct {
	k8sClient *k8s.KubernetesManager
}

func (e Executer) Run(ctx context.Context, queueObj queue_models.QueueObj) {
	jobTemplate, err := e.LoadJobTemplate(queueObj.Environment.Link)
	if err != nil {
		log.Printf("Error loading Job template: %v\n", err)
		// Handle the error according to your application's requirements
		return
	}

	// Create a Kubernetes Job using the parsed template
	err = e.CreateJobFromTemplate(queueObj.Submission.Link, queueObj.TestData.Link, jobTemplate)
	if err != nil {
		log.Printf("Error creating Kubernetes Job: %v\n", err)
		// Handle the error according to your application's requirements
	}
}

func NewExecuter(client *k8s.KubernetesManager) Executer {
	return Executer{
		k8sClient: client,
	}
}

func (e Executer) LoadJobTemplate(configMapName string) ([]byte, error) {
	// Retrieve ConfigMap data
	configMap, err := e.k8sClient.Clientset.CoreV1().ConfigMaps(e.k8sClient.Namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// Assuming the Job template is stored in a key named "job-template"
	for _, val := range configMap.BinaryData {
		return val, nil
	}

	return nil, fmt.Errorf("Job template not found in ConfigMap")
}

func (e Executer) CreateJobFromTemplate(jobName string, testName string, jobTemplate []byte) error {
	// Parse the YAML template into a v1.Job object
	parsedJob, err := utils.ParseJobTemplate(jobTemplate)
	if err != nil {
		return err
	}

	// Set the name for the new Job
	parsedJob.ObjectMeta.Name = jobName
	parsedJob.Spec.Template.Spec.Volumes = append(parsedJob.Spec.Template.Spec.Volumes, corev1.Volume{
		Name: "solution",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: jobName},
			},
		},
	}, corev1.Volume{
		Name: "test",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: testName},
			},
		},
	})

	container := &parsedJob.Spec.Template.Spec.Containers[0]

	// Add volume mounts for the specified paths in the container.
	container.VolumeMounts = append(container.VolumeMounts,
		corev1.VolumeMount{
			Name:      "solution",
			MountPath: "/path/to/solution",
		},
		corev1.VolumeMount{
			Name:      "test",
			MountPath: "/path/to/test",
		},
	)

	// Create the Job using the parsed template
	_, err = e.k8sClient.Clientset.BatchV1().Jobs(e.k8sClient.Namespace).Create(context.TODO(), parsedJob, metav1.CreateOptions{})
	return err
}
