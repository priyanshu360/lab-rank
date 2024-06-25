package postgres

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type editorPostgres struct {
	db *gorm.DB
}

// NewEditorPostgresRepo creates a new PostgreSQL repository for editors.
func NewEditorPostgresRepo(db *gorm.DB) *editorPostgres {
	if err := db.AutoMigrate(&models.Editor{}); err != nil {
		panic(err)
	}
	return &editorPostgres{db}
}

// CreateEditor creates a new editor.
func (psql *editorPostgres) CreateEditor(ctx context.Context, editor *models.Editor) models.AppError {
	editor.CreatedAt = time.Now()
	editor.LastUpdated = time.Now()
	result := psql.db.WithContext(ctx).Create(editor)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	log.Println(editor, result)
	return models.NoError
}

// GetEditorByID retrieves an editor by their ID.
func (psql *editorPostgres) GetEditorByID(ctx context.Context, editorID int) (models.Editor, models.AppError) {
	var editor models.Editor
	result := psql.db.WithContext(ctx).First(&editor, editorID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Editor not found
			return editor, models.EditorNotFoundError
		}
		return editor, models.InternalError.Add(result.Error)
	}
	return editor, models.NoError
}

// GetEditorByUserID retrieves editors by user ID.
func (psql *editorPostgres) GetEditorByUserID(ctx context.Context, userID uuid.UUID) ([]models.Editor, models.AppError) {
	var editors []models.Editor
	result := psql.db.WithContext(ctx).Where("user_id = ?", userID).Find(&editors)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}
	return editors, models.NoError
}

func (psql *editorPostgres) GetEditorByUserIDAndProblemID(ctx context.Context, userID uuid.UUID, problemID int) (models.Editor, models.AppError) {
	var editor models.Editor
	result := psql.db.WithContext(ctx).Where("user_id = ? AND problem_id = ?", userID, problemID).First(&editor)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return editor, models.EditorNotFoundError
		}
		return editor, models.InternalError.Add(result.Error)
	}
	return editor, models.NoError
}

// UpdateEditor updates an editor's information.
func (psql *editorPostgres) UpdateEditor(ctx context.Context, editorID int, editor models.Editor) models.AppError {
	// Check if the editor with the provided ID exists before updating
	var existingEditor models.Editor
	result := psql.db.WithContext(ctx).First(&existingEditor, editorID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Editor not found
			return models.EditorNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the update
	editor.LastUpdated = time.Now()
	result = psql.db.WithContext(ctx).Model(&existingEditor).Updates(editor)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// DeleteEditor deletes an editor by their ID.
func (psql *editorPostgres) DeleteEditor(ctx context.Context, editorID int) models.AppError {
	// Check if the editor with the provided ID exists before deletion
	var existingEditor models.Editor
	result := psql.db.WithContext(ctx).First(&existingEditor, editorID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Editor not found
			return models.EditorNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the deletion
	result = psql.db.WithContext(ctx).Delete(&existingEditor)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// ListEditors lists editors with pagination.
func (psql *editorPostgres) ListEditors(ctx context.Context, page int, pageSize int) ([]models.Editor, models.AppError) {
	var editors []models.Editor
	result := psql.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&editors)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}
	return editors, models.NoError
}

// ListEditorsWithLimit lists editors with pagination, returning pointers to the editor objects.
func (psql *editorPostgres) ListEditorsWithLimit(ctx context.Context, page int, pageSize int) ([]*models.Editor, models.AppError) {
	var editors []*models.Editor

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch editors with the specified pagination
	result := psql.db.Offset(offset).Limit(pageSize).Find(&editors)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return editors, models.NoError
}
