package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	batchv1 "k8s.io/api/batch/v1"
)

func TestParseJobTemplate(t *testing.T) {
	yamlTemplate := `
apiVersion: batch/v1
kind: Job
metadata:
  name: example-job
spec:
  template:
    spec:
      containers:
      - name: example-container
        image: busybox
        command: ["echo", "Hello, Kubernetes!"]
  backoffLimit: 4
`

	job, err := ParseJobTemplate([]byte(yamlTemplate))

	assert.Nil(t, err, "Unexpected error during parsing")
	assert.IsType(t, &batchv1.Job{}, job, "Unexpected type returned")
	assert.Equal(t, "example-job", job.ObjectMeta.Name, "Unexpected Job name")
	assert.Len(t, job.Spec.Template.Spec.Containers, 1, "Unexpected number of containers")
	assert.Equal(t, "busybox", job.Spec.Template.Spec.Containers[0].Image, "Unexpected container image")
	assert.Equal(t, int32(4), *job.Spec.BackoffLimit, "Unexpected backoff limit")
}

// func TestParseJobTemplateWithInvalidYAML(t *testing.T) {
// 	// Example of invalid YAML template
// 	invalidYAMLTemplate := `invalid yaml`

// 	// Run the ParseJobTemplate function with invalid YAML
// 	job, err := ParseJobTemplate(invalidYAMLTemplate)

// 	// Check for errors
// 	assert.NotNil(t, err, "Expected an error for invalid YAML")
// 	assert.Nil(t, job, "Expected nil Job object when there is an error")
// }
