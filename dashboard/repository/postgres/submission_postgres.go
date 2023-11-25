package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
	"gorm.io/gorm"
)

type submissionPostgres struct {
	db *gorm.DB
}

// NewSubmissionPostgresRepo creates a new PostgreSQL repository for submissions.
func NewSubmissionPostgresRepo(db *gorm.DB) *submissionPostgres {
	return &submissionPostgres{db}
}

// CreateSubmission creates a new submission.
func (psql *submissionPostgres) CreateSubmission(ctx context.Context, submission models.Submission) models.AppError {
	result := psql.db.WithContext(ctx).Create(submission)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSubmissionByID retrieves a submission by its ID.
func (psql *submissionPostgres) GetSubmissionByID(ctx context.Context, submissionID uuid.UUID) (models.Submission, models.AppError) {
	var submission models.Submission
	result := psql.db.WithContext(ctx).First(&submission, submissionID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Submission not found
			return submission, models.SubmissionNotFoundError
		}
		return submission, models.InternalError.Add(result.Error)
	}
	return submission, models.NoError
}

func (psql *submissionPostgres) GetSubmissionsListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Submission, models.AppError) {
	var submissions []*models.Submission

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch submissions with the specified pagination
	result := psql.db.Offset(offset).Table("lab_rank.submissions").Limit(pageSize).Find(&submissions)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return submissions, models.NoError
}

// Add other repository methods for submissions as needed.

func (psql *submissionPostgres) GetQueueData(ctx context.Context, submission models.Submission) (queue_models.QueueObj, models.AppError) {
	log.Println(submission.ProblemID)
	var envArray models.EnvironmentJSON
	if err := psql.db.Model(&models.Problem{}).
		Select("environment").
		Where("id = ?", submission.ProblemID).
		First(&envArray).Error; err != nil {
		return queue_models.QueueObj{}, models.InternalError.Add(err)
	}

	var environmentID uuid.UUID
	var environmentLink string
	for _, env := range envArray {
		if env.Language == submission.Lang {
			environmentID = env.Id
		}
	}

	if err := psql.db.Model(&models.Environment{}).
		Select("link").Table("environment").
		Where("id = ?", environmentID).
		First(&environmentLink).Error; err != nil {
		return queue_models.QueueObj{}, models.InternalError.Add(err)
	}

	var testData models.TestLinkJSON
	if err := psql.db.Model(&models.Problem{}).
		Select("test_links").
		Where("id = ?", submission.ProblemID).
		First(&testData).Error; err != nil {
		return queue_models.QueueObj{}, models.InternalError.Add(err)
	}

	var testLink string
	for _, test := range testData {
		if test.Language == submission.Lang {
			testLink = test.Link
		}
	}

	queue := *queue_models.NewQueueObj(submission.ID, submission.Link, environmentID, environmentLink, testLink)

	return queue, models.NoError
}
