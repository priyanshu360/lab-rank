package executer

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"github.com/priyanshu360/lab-rank/judge/models"
// 	"github.com/priyanshu360/lab-rank/judge/service/k8s"
// 	"k8s.io/client-go/util/homedir"
// )

// func TestRun(t *testing.T) {
// 	// Create a context with a timeout to avoid infinite loops in case of errors
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Create an instance of the Executer using the MockExecuter for testing

// 	kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "path to kubeconfig file")
// 	namespace := flag.String("namespace", "storage", "Kubernetes namespace")
// 	flag.Parse()

// 	k8s, err := k8s.NewKubernetesManager(*kubeconfig, *namespace)
// 	if err != nil {
// 		fmt.Printf("Error creating KubernetesManager: %v\n", err)
// 		os.Exit(1)
// 	}
// 	executer := NewExecuter(k8s)

// 	// Mock the submission data
// 	submission := models.SubmissionData{
// 		Link: "environment-88dbed99-321f-4918-af57-b641831dda84",
// 		// Add other fields as needed for your actual data structure
// 	}

// 	// Run the method under test
// 	executer.Run(ctx, submission, "environment-08ab6b51-b0bc-4d5f-afc1-51495f8a398f")

// 	// Add assertions based on the expected behavior
// 	// For example, assert that the job was created successfully
// 	// assert.Nil(t, executer.CreateJobFromTemplate("example-job", "mock-template"))

// 	// If LoadJobTemplate or CreateJobFromTemplate return an error, you can assert the error as well
// 	// For example:
// 	// assert.Nil(t, executer.LoadJobTemplate("example-link"))
// 	// assert.NotNil(t, executer.CreateJobFromTemplate("example-job", "mock-template"))
// }
