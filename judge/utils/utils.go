package utils

import (
	"fmt"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func ParseJobTemplate(template []byte) (*batchv1.Job, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(template, nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	job := obj.(*batchv1.Job)

	log.Print(job)

	return job, nil
}
