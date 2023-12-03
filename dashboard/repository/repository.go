package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
)

type UserRepository interface {
	CreateUser(context.Context, models.User) models.AppError
	GetUserByID(context.Context, uuid.UUID) (models.User, models.AppError)
	GetUserByEmail(context.Context, string) (models.User, models.AppError)
	UpdateUser(context.Context, uuid.UUID, models.User) models.AppError
	DeleteUser(context.Context, uuid.UUID) models.AppError
	ListUsers(context.Context, int, int) ([]models.User, models.AppError)
	ListUsersWisthLimit(context.Context, int, int) ([]*models.User, models.AppError)
}

type UniversityRepository interface {
	CreateUniversity(context.Context, models.University) models.AppError
	GetUniversityByID(context.Context, uuid.UUID) (models.University, models.AppError)
	GetUniversitiesListByLimit(context.Context, int, int) ([]*models.University, models.AppError)
	UpdateUniversity(context.Context, uuid.UUID, models.University) models.AppError
	// Add other repository methods specific to University
}

type SyllabusRepository interface {
	CreateSyllabus(context.Context, models.Syllabus) models.AppError
	GetSyllabusByID(context.Context, uuid.UUID) (models.Syllabus, models.AppError)
	GetSyllabusListByLimit(context.Context, int, int) ([]*models.Syllabus, models.AppError)
	UpdateSyllabus(context.Context, uuid.UUID, models.Syllabus) models.AppError
	// Add other repository methods specific to Syllabus
}

type SubjectRepository interface {
	CreateSubject(context.Context, models.Subject) models.AppError
	GetSubjectByID(context.Context, uuid.UUID) (models.Subject, models.AppError)
	GetSubjectsListByLimit(context.Context, int, int) ([]*models.Subject, models.AppError)
	UpdateSubject(context.Context, uuid.UUID, models.Subject) models.AppError
	// Add other repository methods specific to Subject
}

type ProblemRepository interface {
	CreateProblem(context.Context, models.Problem) models.AppError
	GetProblemByID(context.Context, uuid.UUID) (models.Problem, models.AppError)
	GetProblemsListByLimit(context.Context, int, int) ([]*models.Problem, models.AppError)
	UpdateProblem(context.Context, uuid.UUID, models.Problem) models.AppError
	// Add other repository methods specific to Problem
}

type EnvironmentRepository interface {
	CreateEnvironment(context.Context, models.Environment) models.AppError
	GetEnvironmentByID(context.Context, uuid.UUID) (models.Environment, models.AppError)
	GetEnvironmentsListByLimit(context.Context, int, int) ([]*models.Environment, models.AppError)
	UpdateEnvironment(context.Context, uuid.UUID, models.Environment) models.AppError
	// Add other repository methods specific to Environment
}

type SubmissionRepository interface {
	CreateSubmission(context.Context, models.Submission) models.AppError
	GetSubmissionByID(context.Context, uuid.UUID) (models.Submission, models.AppError)
	GetSubmissionsListByLimit(context.Context, int, int) ([]*models.Submission, models.AppError)
	GetQueueData(context.Context, models.Submission) (queue_models.QueueObj, models.AppError)
	UpdateSubmission(context.Context, uuid.UUID, models.Submission) models.AppError
	// Add other repository methods specific to Submission
}

type CollegeRepository interface {
	CreateCollege(context.Context, models.College) models.AppError
	GetCollegeByID(context.Context, uuid.UUID) (models.College, models.AppError)
	GetCollegesListByLimit(context.Context, int, int) ([]*models.College, models.AppError)
	UpdateCollege(context.Context, uuid.UUID, models.College) models.AppError
	// Add other repository methods specific to College
}

type FileSystem interface {
	StoreFile(context.Context, []byte, uuid.UUID, models.FileType, string) (string, models.AppError)
	GetFile(context.Context, string) ([]byte, models.AppError)
	// MakeFileName(...string) string
}
