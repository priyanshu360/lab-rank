package problem

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

const LanguageExtension = "py"

func (s *problemService) storeFile(ctx context.Context, file []byte, fileType string, fileId uuid.UUID, title string) (uuid.UUID, models.AppError) {
	fileID := fileId
	// Specify the directory where files will be stored
	storageDir := fmt.Sprintf("./uploads/%s", fileID)
	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return fileID, models.InternalError.Add(err)
	}

	// Specify the file path based on the file ID and type
	filePath := filepath.Join(storageDir, fmt.Sprintf("%s.%s", title, LanguageExtension))
	log.Println(filePath)
	// Write the file content to the specified path
	data, err := url.QueryUnescape(string(file))
	if err != nil {
		return fileID, models.InternalError.Add(err)
	}
	if err = os.WriteFile(filePath, []byte(data), os.ModePerm); err != nil {
		return fileID, models.InternalError.Add(err)
	}

	return fileID, models.NoError
}

// generateFileLink is a helper function to generate a file link based on the file ID
func generateFileLink(fileID uuid.UUID, fileType string) string {
	// Assuming you have a route in your API to serve files, e.g., "/files/{fileType}/{fileID}"
	return fmt.Sprintf("/files/%s/%s", fileType, fileID)
}
