package postgres

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type syllabusPostgres struct {
	db *gorm.DB
}

// NewSyllabusPostgresRepo creates a new PostgreSQL repository for syllabi.
func NewSyllabusPostgresRepo(db *gorm.DB) *syllabusPostgres {
	if err := db.AutoMigrate(models.Syllabus{}); err != nil {
		panic(err)
	}
	return &syllabusPostgres{db}
}

// CreateSyllabus creates a new syllabus.
func (psql *syllabusPostgres) CreateSyllabus(ctx context.Context, syllabus models.Syllabus) models.AppError {
	result := psql.db.WithContext(ctx).Create(&syllabus)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSyllabusByID retrieves a syllabus by its ID.
func (psql *syllabusPostgres) GetSyllabusByID(ctx context.Context, syllabusID int) (models.Syllabus, models.AppError) {
	var syllabus models.Syllabus
	result := psql.db.WithContext(ctx).First(&syllabus, syllabusID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Syllabus not found
			return syllabus, models.SyllabusNotFoundError
		}
		return syllabus, models.InternalError.Add(result.Error)
	}
	return syllabus, models.NoError
}

func (psql *syllabusPostgres) GetSyllabusListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Syllabus, models.AppError) {
	var syllabi []*models.Syllabus

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch syllabi with the specified pagination
	result := psql.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&syllabi)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return syllabi, models.NoError
}

// Add other repository methods for syllabi as needed.

func (psql *syllabusPostgres) GetCollegeIDsForUniversityID(ctx context.Context, universityID int) ([]int, models.AppError) {
	var collegeIDs []int // Assuming college_id is of type int, adjust accordingly

	result := psql.db.WithContext(ctx).Where("university_id = ?", universityID).
		Pluck("id", &collegeIDs)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// College not found
			return collegeIDs, models.CollegeNotFoundError
		}
		return collegeIDs, models.InternalError.Add(result.Error)
	}
	return collegeIDs, models.NoError
}

func (psql *syllabusPostgres) GetSubjectsByUniversityID(ctx context.Context, universityID int) ([]models.Subject, models.AppError) {
	var subjects []models.Subject
	result := psql.db.WithContext(ctx).Where("university_id = ?", universityID).Find(&subjects)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Subject not found
			return subjects, models.SubjectNotFoundError
		}
		return subjects, models.InternalError.Add(result.Error)
	}
	return subjects, models.NoError
}

// GetSyllabusBySubjectID retrieves a syllabus by its subject ID.
func (psql *syllabusPostgres) GetSyllabusBySubjectID(ctx context.Context, subjectID int) ([]*models.Syllabus, models.AppError) {
	var syllabi []*models.Syllabus
	result := psql.db.WithContext(ctx).Where("subject_id = ?", subjectID).Find(&syllabi)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Syllabus not found
			return syllabi, models.SyllabusNotFoundError
		}
		return syllabi, models.InternalError.Add(result.Error)
	}
	return syllabi, models.NoError
}

// func (psql *syllabusPostgres) UpdateUserAccessIDs(ctx context.Context, user models.User) models.AppError {
// 	// Step 1: Get user's university ID using college ID
// 	var universityID string
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.college").
// 		Where("id = ?", user.CollegeID).
// 		Pluck("university_id", &universityID).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	// Step 2: Get all syllabus IDs for the university
// 	var syllabusIDs []string
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.syllabus").
// 		Where("uni_college_id = ? AND syllabus_level = ?", universityID, models.SyllabusLevelCollege).
// 		Pluck("id", &syllabusIDs).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	var collegeSyllabusIDs []string
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.syllabus").
// 		Where("uni_college_id = ? AND syllabus_level = ?", user.CollegeID, models.SyllabusLevelCollege).
// 		Pluck("id", &collegeSyllabusIDs).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	syllabusIDs = append(syllabusIDs, collegeSyllabusIDs...)

// 	// Step 3: Get all access IDs for the syllabuses
// 	var accessIDs []int
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.access_level").
// 		Where("syllabus_id IN (?) AND mode = STUDENT", syllabusIDs).
// 		Pluck("id", &accessIDs).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	// Step 4: Update auth table entry for the user with new access IDs
// 	var auth models.Auth
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.auth").
// 		Where("user_id = ?", user.ID).
// 		First(&auth).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	// Add new access IDs to the existing ones

// 	// Save the updated entry
// 	if err := psql.db.WithContext(ctx).Table("lab_rank.auth").
// 		Where("user_id = ?", user.ID).
// 		Save(&auth).
// 		Error; err != nil {
// 		return models.InternalError.Add(err)
// 	}

// 	return models.NoError
// }
