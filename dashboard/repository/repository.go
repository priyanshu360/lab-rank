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
	CreateUniversity(context.Context, *models.University) models.AppError
	GetUniversityByID(context.Context, int) (models.University, models.AppError)
	GetUniversitiesListByLimit(context.Context, int, int) ([]*models.University, models.AppError)
	GetAllUniversityNames(context.Context) ([]*models.UniversityIdName, models.AppError)
	// Add other repository methods specific to University
}

type SyllabusRepository interface {
	CreateSyllabus(context.Context, models.Syllabus) models.AppError
	GetSyllabusByID(context.Context, int) (models.Syllabus, models.AppError)
	GetSyllabusListByLimit(context.Context, int, int) ([]*models.Syllabus, models.AppError)
	GetCollegeIDsForUniversityID(context.Context, int) ([]int, models.AppError)
	GetSubjectsByUniversityID(context.Context, int) ([]models.Subject, models.AppError)
	GetSyllabusBySubjectID(context.Context, int) ([]*models.Syllabus, models.AppError)

	// UpdateUserAccessIDs(ctx context.Context, user models.User) models.AppError
	// Add other repository methods specific to Syllabus
}

type SubjectRepository interface {
	CreateSubject(context.Context, models.Subject) models.AppError

	GetSubjectsByUniversityID(context.Context, int) ([]*models.Subject, models.AppError)
	GetSubjectByID(context.Context, int) (models.Subject, models.AppError)
	GetSubjectsListByLimit(context.Context, int, int) ([]*models.Subject, models.AppError)
	// Add other repository methods specific to Subject
}

type ProblemRepository interface {
	CreateProblem(context.Context, *models.Problem) models.AppError
	GetProblemByID(context.Context, int) (models.Problem, models.AppError)
	GetProblemsListByLimit(context.Context, int, int) ([]*models.Problem, models.AppError)
	GetProblemsForSubject(context.Context, int, int) ([]*models.Problem, models.AppError)
	// Add other repository methods specific to Problem
}

type EnvironmentRepository interface {
	CreateEnvironment(context.Context, models.Environment) models.AppError
	GetEnvironmentByID(context.Context, int) (models.Environment, models.AppError)
	GetEnvironmentsListByLimit(context.Context, int, int) ([]*models.Environment, models.AppError)
	// Add other repository methods specific to Environment
}

type SubmissionRepository interface {
	CreateSubmission(context.Context, models.Submission) models.AppError
	GetSubmissionByID(context.Context, int) (models.Submission, models.AppError)
	GetSubmissionsListByLimit(context.Context, int, int) ([]*models.Submission, models.AppError)
	GetQueueData(context.Context, models.Submission) (queue_models.QueueObj, models.AppError)
	UpdateSubmission(context.Context, int, models.Submission) models.AppError
	GetSubmissionsByUserID(context.Context, uuid.UUID) ([]*models.Submission, models.AppError)
	GetSubmissionsWithTitleByUserID(context.Context, uuid.UUID) ([]*models.SubmissionWithProblemTitle, models.AppError)
	// Add other repository methods specific to Submission
}

type CollegeRepository interface {
	CreateCollege(context.Context, models.College) models.AppError
	GetCollegeByID(context.Context, int) (models.College, models.AppError)
	GetCollegesListByLimit(context.Context, int, int) ([]*models.College, models.AppError)
	GetCollegesByUniversityID(context.Context, int) ([]*models.CollegeIdName, models.AppError)
	// Add other repository methods specific to College
}

type FileSystem interface {
	StoreFile(context.Context, []byte, string, models.FileType, string) (string, models.AppError)
	GetFile(context.Context, string) ([]byte, models.AppError)
	// MakeFileName(...string) string
}

type AuthRepository interface {
	SignUp(context.Context, models.User, models.Auth) models.AppError
	GetUserAuthByEmail(context.Context, string) (*models.User, *models.Auth, models.AppError)
}

type SessionRepository interface {
	GetSession(context.Context, uuid.UUID) (*models.AuthSession, models.AppError)
	SetSession(context.Context, *models.AuthSession) (uuid.UUID, models.AppError)
	RemoveSession(context.Context, uuid.UUID) models.AppError
}

type EditorRepository interface {
	CreateEditor(context.Context, *models.Editor) models.AppError
	GetEditorByID(context.Context, int) (models.Editor, models.AppError)
	GetEditorByUserID(context.Context, uuid.UUID) ([]models.Editor, models.AppError)
	GetEditorByUserIDAndProblemID(context.Context, uuid.UUID, int) (models.Editor, models.AppError)
	UpdateEditor(context.Context, int, models.Editor) models.AppError
	DeleteEditor(context.Context, int) models.AppError
	ListEditors(context.Context, int, int) ([]models.Editor, models.AppError)
	ListEditorsWithLimit(context.Context, int, int) ([]*models.Editor, models.AppError)
}
