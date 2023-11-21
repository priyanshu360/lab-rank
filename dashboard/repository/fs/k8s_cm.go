package filesys

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type k8sCMStore struct {
	clientset *kubernetes.Clientset
	namespace string
}

func NewK8sCMStore(clientset *kubernetes.Clientset, namespace string) *k8sCMStore {
	return &k8sCMStore{
		clientset: clientset,
		namespace: namespace,
	}
}

func (fs *k8sCMStore) StoreFile(ctx context.Context, content []byte, id uuid.UUID, ftype models.FileType, extension string) (string, models.AppError) {
	configMapName := fmt.Sprintf("%s-%s", string(ftype), id.String())
	configMapData := map[string][]byte{
		"file": content,
	}

	cm := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: fs.namespace,
		},
		BinaryData: configMapData,
	}

	_, err := fs.clientset.CoreV1().ConfigMaps(fs.namespace).Create(context.TODO(), cm, metav1.CreateOptions{})
	if err != nil {
		return "", models.InternalError.Add(err)
	}

	return configMapName, models.NoError
}

func (fs *k8sCMStore) GetFile(ctx context.Context, configMapName string) ([]byte, models.AppError) {
	cm, err := fs.clientset.CoreV1().ConfigMaps(fs.namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, models.InternalError.Add(err)
	}

	content, found := cm.BinaryData["file"]
	if !found {
		return nil, models.InternalError.Add(fmt.Errorf("File not found"))
	}

	return []byte(content), models.NoError
}
