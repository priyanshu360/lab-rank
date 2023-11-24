package executer

import (
	"context"
	"fmt"
	"log"

	"github.com/priyanshu360/lab-rank/judge/models"
	"github.com/priyanshu360/lab-rank/judge/service/k8s"
	"github.com/priyanshu360/lab-rank/judge/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Executer struct {
	k8sClient *k8s.KubernetesManager
}

func (e Executer) Run(ctx context.Context, submission models.SubmissionData, environmentLink string) {
	log.Println(submission)

	// Load Job template from the ConfigMap
	jobTemplate, err := e.LoadJobTemplate(environmentLink)
	if err != nil {
		log.Printf("Error loading Job template: %v\n", err)
		// Handle the error according to your application's requirements
		return
	}

	// Create a Kubernetes Job using the parsed template
	err = e.CreateJobFromTemplate(submission.Link, jobTemplate)
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
	jobTemplate, found := configMap.BinaryData["file"]
	if !found {
		return nil, fmt.Errorf("Job template not found in ConfigMap")
	}

	return jobTemplate, nil
}

func (e Executer) CreateJobFromTemplate(jobName string, jobTemplate []byte) error {
	// Parse the YAML template into a v1.Job object
	parsedJob, err := utils.ParseJobTemplate(jobTemplate)
	if err != nil {
		return err
	}

	// Set the name for the new Job
	parsedJob.ObjectMeta.Name = jobName
	parsedJob.Spec.Template.Spec.Volumes = append(parsedJob.Spec.Template.Spec.Volumes, corev1.Volume{
		Name: "Solution",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: jobName},
			},
		},
	})

	// Create the Job using the parsed template
	_, err = e.k8sClient.Clientset.BatchV1().Jobs(e.k8sClient.Namespace).Create(context.TODO(), parsedJob, metav1.CreateOptions{})
	return err
}
