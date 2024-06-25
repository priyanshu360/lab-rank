package editor

import (
	"context"
	"fmt"
	"time"

	cfg "github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service interface {
	Start(context.Context, *models.Editor) (models.Editor, models.AppError)
}

type service struct {
	editor      repository.EditorRepository
	environment repository.EnvironmentRepository
}

func New(editor repository.EditorRepository, environment repository.EnvironmentRepository) *service {
	return &service{
		editor:      editor,
		environment: environment,
	}
}

func (svc *service) Start(ctx context.Context, editor *models.Editor) (models.Editor, models.AppError) {
	// find if editor already exist
	// validate editor state and take action

	oEditor, err := svc.editor.GetEditorByUserIDAndProblemID(ctx, editor.UserID, editor.ProblemID)
	if err != models.NoError && err != models.EditorNotFoundError {
		return *editor, err
	}

	if err == models.EditorNotFoundError {
		err = svc.editor.CreateEditor(ctx, editor)
		if err != models.NoError {
			return *editor, err
		}
	} else {
		editor = &oEditor
	}

	_, err = svc.environment.GetEnvironmentByID(ctx, editor.Environment.Id)
	if err != models.NoError {
		return *editor, err
	}

	image := "lscr.io/linuxserver/code-server:latest"
	pods := 1
	ports := []int{8080}

	if editor.Deployment == "" {
		editor.Deployment = fmt.Sprintf("editor-pid-%d-%d", editor.ProblemID, editor.ID)
	}

	_, kerr := createDeployment(editor.Deployment, image, pods, ports)
	if kerr != nil {
		return *editor, models.InternalError.Add(kerr)
	}

	editor.Status = models.Scaled

	err = svc.editor.UpdateEditor(ctx, editor.ID, *editor)
	return *editor, err
}

func createDeployment(name, image string, pods int, ports []int) (podIPs []string, err error) {
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	return nil, fmt.Errorf("error creating in-cluster config: %v", err)
	// }

	config, err := clientcmd.BuildConfigFromFlags("", cfg.K8sConfig)
	if err != nil {
		return podIPs, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating Kubernetes client: %v", err)
	}

	var containerPorts []corev1.ContainerPort
	for _, port := range ports {
		containerPorts = append(containerPorts, corev1.ContainerPort{
			ContainerPort: int32(port),
			Protocol:      corev1.ProtocolTCP,
		})
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(int32(pods)),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "editor",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "editor",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "editor-container",
							Image: image,
							Ports: containerPorts,
						},
					},
				},
			},
		},
	}

	deploymentsClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error creating deployment: %v", err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	podIPs, err = waitForPods(clientset, "editor", pods)
	if err != nil {
		return nil, fmt.Errorf("error waiting for pods to be ready: %v", err)
	}

	return podIPs, nil
}

func int32Ptr(i int32) *int32 { return &i }

func waitForPods(clientset *kubernetes.Clientset, appLabel string, numPods int) ([]string, error) {
	timeout := time.After(5 * time.Minute)
	tick := time.Tick(2 * time.Second)

	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("timed out waiting for pods to be ready")
		case <-tick:
			pods, err := clientset.CoreV1().Pods(corev1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{
				LabelSelector: fmt.Sprintf("app=%s", appLabel),
			})
			if err != nil {
				return nil, fmt.Errorf("error listing pods: %v", err)
			}

			var readyPodIPs []string
			for _, pod := range pods.Items {
				if pod.Status.Phase == corev1.PodRunning {
					for _, cond := range pod.Status.Conditions {
						if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
							readyPodIPs = append(readyPodIPs, pod.Status.PodIP)
						}
					}
				}
			}

			if len(readyPodIPs) >= numPods {
				return readyPodIPs, nil
			}
		}
	}
}
