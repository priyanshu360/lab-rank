// submission_repository.go

package postgres

import (
	"context"

	"github.com/google/uuid"
	dashboard_models "github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/judge/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{db}
}

func (r *SubmissionRepository) GetNextBatch(ctx context.Context, status dashboard_models.Status, batchSize int) []dashboard_models.Submission {
	var submissions []dashboard_models.Submission
	r.db.WithContext(ctx).
		Where("status = ?", status).
		Limit(batchSize).
		Find(&submissions)
	return submissions
}

func (r *SubmissionRepository) Update(ctx context.Context, submission models.SubmissionData) {
	r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"status"}),
		}).
		Create(&submission)
}

func (r *SubmissionRepository) GetEnvFromPIDAndLang(ctx context.Context, problemID uuid.UUID, lang dashboard_models.ProgrammingLanguageEnum) dashboard_models.Environment {
	var environments []dashboard_models.Environment
	r.db.WithContext(ctx).
		Where("problem_id = ? AND lang = ?", problemID, lang).
		Find(&environments)
	return environments[0]
}
