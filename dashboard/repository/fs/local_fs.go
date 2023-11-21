package filesys

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type local struct {
	basePath string
}

func NewLocalFS(path string) local {
	return local{
		basePath: path,
	}
}

func (fs local) StoreFile(ctx context.Context, content []byte, id uuid.UUID, ftype models.FileType, extension string) (string, models.AppError) {
	typePath := filepath.Join(fs.basePath, string(ftype))
	if err := os.MkdirAll(typePath, 0755); err != nil {
		return "", models.InternalError.Add(err)
	}

	filename := fmt.Sprintf("%s.%s", id.String(), extension)
	filePath := filepath.Join(typePath, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return "", models.InternalError.Add(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return "", models.InternalError.Add(err)
	}

	return filePath, models.NoError
}

func (fs local) GetFile(ctx context.Context, filePath string) ([]byte, models.AppError) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, models.InternalError.Add(fmt.Errorf("File not found %s", err.Error()))
		}
		return nil, models.InternalError.Add(err)
	}

	return content, models.NoError
}
