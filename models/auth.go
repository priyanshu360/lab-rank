package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type AccessIDs []uuid.UUID

// AccessLevel represents the lab_rank.access_level table in the database.
type AccessLevelModeEnum string

const (
	AccessLevelAdmin   AccessLevelModeEnum = "ADMIN"
	AccessLevelTeacher AccessLevelModeEnum = "TEACHER"
	AccessLevelStudent AccessLevelModeEnum = "STUDENT"
	AccessLevelNone    AccessLevelModeEnum = "NONE"
)

var CanAccess map[AccessLevelModeEnum]map[AccessLevelModeEnum]bool = map[AccessLevelModeEnum]map[AccessLevelModeEnum]bool{
	AccessLevelAdmin: {
		AccessLevelAdmin:   true,
		AccessLevelTeacher: true,
		AccessLevelStudent: true,
		AccessLevelNone:    true,
	},
	AccessLevelTeacher: {
		AccessLevelAdmin:   false,
		AccessLevelTeacher: true,
		AccessLevelStudent: true,
		AccessLevelNone:    true,
	},
	AccessLevelStudent: {
		AccessLevelAdmin:   false,
		AccessLevelTeacher: false,
		AccessLevelStudent: true,
		AccessLevelNone:    true,
	},
	AccessLevelNone: {
		AccessLevelAdmin:   false,
		AccessLevelTeacher: false,
		AccessLevelStudent: false,
		AccessLevelNone:    false,
	},
}

type Auth struct {
	UserID       uuid.UUID           `json:"user_id"`
	Salt         []byte              `json:"salt"`
	PasswordHash string              `json:"password_hash"`
	Mode         AccessLevelModeEnum `json:"mode"`
}

type AuthSession struct {
	User
	Mode AccessLevelModeEnum `json:"mode"`
}

func NewAuthSession(user *User, mode AccessLevelModeEnum) *AuthSession {
	return &AuthSession{
		User: *user,
		Mode: mode,
	}
}

type AuthenticateAPIResponse struct {
	Message *AuthSession
}

func NewAuthenticateAPIResponse(jwt *AuthSession) *AuthenticateAPIResponse {
	return &AuthenticateAPIResponse{
		Message: jwt,
	}
}

type SignUpAPIRequest struct {
	CreateUserAPIRequest
	Password string `json:"password"`
}

type SignUpAPIResponse CreateUserAPIResponse

type LoginAPIResponse struct {
	Message string
}

func NewLoginAPIResponse(jwt string) *LoginAPIResponse {
	return &LoginAPIResponse{
		Message: jwt,
	}
}

func (res *AuthenticateAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

func (res *LoginAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

type LoginAPIRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (r SignUpAPIRequest) ToUser() *User {
	return &User{
		ID:           uuid.New(),
		CollegeID:    r.CollegeID,
		Status:       UserStatusActive, // You can set an initial status here if needed
		Email:        r.Email,
		ContactNo:    r.ContactNo,
		UniversityID: r.UniversityID,
		DOB:          r.DOB,
		Name:         r.Name,
		UserName:     r.UserName,
	}
}

// Value implements the driver.Valuer interface
func (e AccessIDs) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// Scan implements the sql.Scanner interface
func (e *AccessIDs) Scan(value interface{}) error {
	if value == nil {
		*e = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var ids []uuid.UUID
		if err := json.Unmarshal(v, &ids); err != nil {
			return err
		}
		*e = AccessIDs(ids)
		return nil
	default:
		return errors.New("unsupported type for AccessIDs")
	}
}

// College struct
type College struct {
	ID           uuid.UUID       `json:"id" validate:"required"`
	Title        string          `json:"title" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
}

// CreateCollegeAPIRequest struct
type CreateCollegeAPIRequest struct {
	Title        string          `json:"title" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
}

// CollegeAPIResponse struct
type CollegeAPIResponse struct {
	Message *College
}

// Implement the Parse method for POST request for CreateCollegeAPIRequest
func (r *CreateCollegeAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for CollegeAPIResponse
func (cr *CollegeAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}

func (r *CreateCollegeAPIRequest) ToCollege() *College {
	return &College{
		ID:           uuid.New(),
		Title:        r.Title,
		UniversityID: r.UniversityID,
		Description:  r.Description,
	}
}

func NewCreateCollegeAPIResponse(college *College) *CollegeAPIResponse {
	return &CollegeAPIResponse{
		Message: college,
	}
}

type CollegeIdName struct {
	ID    uuid.UUID
	Title string
}

func NewCollegeIdName(id uuid.UUID, name string) *CollegeIdName {
	return &CollegeIdName{
		ID:    id,
		Title: name,
	}
}

type ListCollegesIdNamesAPIResponse struct {
	Message []*CollegeIdName
}

// Implement the Write method for ListCollegesAPIResponse
func (pr *ListCollegesIdNamesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListCollegesIdNamesAPIResponse(colleges []*CollegeIdName) *ListCollegesIdNamesAPIResponse {
	return &ListCollegesIdNamesAPIResponse{
		Message: colleges,
	}
}

type ListCollegesAPIResponse struct {
	Message []*College
}

// Implement the Write method for ListcollegesAPIResponse
func (pr *ListCollegesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListCollegesAPIResponse(colleges []*College) *ListCollegesAPIResponse {
	return &ListCollegesAPIResponse{
		Message: colleges,
	}
}

// Environment struct
type Environment struct {
	ID             uuid.UUID       `json:"id" validate:"required"`
	Title          string          `json:"title" validate:"required"`
	Link           string          `json:"link" validate:"required"`
	CreatedBy      uuid.UUID       `json:"created_by" validate:"required"`
	CreatedAt      time.Time       `json:"created_at" validate:"required"`
	UpdateEvents   json.RawMessage `json:"update_events" validate:"required"`
	LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids" validate:"required" gorm:"column:live_dockerc_ids"`
	File           []byte          `json:"file" gorm:"-"`
}

// CreateEnvironmentAPIRequest struct
type CreateEnvironmentAPIRequest struct {
	Title          string          `json:"title" validate:"required"`
	CreatedBy      uuid.UUID       `json:"created_by" validate:"required"`
	UpdateEvents   json.RawMessage `json:"update_events"`
	LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids"`
	File           []byte          `json:"file" validate:"required"`
}

// EnvironmentAPIResponse struct
type EnvironmentAPIResponse struct {
	Message *Environment
}

// Implement the Parse method for POST request for CreateEnvironmentAPIRequest
func (r *CreateEnvironmentAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for EnvironmentAPIResponse
func (er *EnvironmentAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(er)
}

func (r *CreateEnvironmentAPIRequest) ToEnvironment() *Environment {
	return &Environment{
		ID:             uuid.New(),
		Title:          r.Title,
		CreatedBy:      r.CreatedBy,
		CreatedAt:      time.Now(),
		UpdateEvents:   []byte("[]"),
		LiveDockerCIDs: []byte("[]"),
		File:           r.File,
	}
}

func NewCreateEnvironmentAPIResponse(environment *Environment) *EnvironmentAPIResponse {
	return &EnvironmentAPIResponse{
		Message: environment,
	}
}

type ListEnvironmentsAPIResponse struct {
	Message []*Environment
}

// Implement the Write method for ListenvironmentsAPIResponse
func (pr *ListEnvironmentsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListEnvironmentsAPIResponse(environments []*Environment) *ListEnvironmentsAPIResponse {
	return &ListEnvironmentsAPIResponse{
		Message: environments,
	}
}

var validate = validator.New()

type AppError struct {
	Type   ErrorType
	Reason string `json:"reason"`
}

func NewAppError(code ErrorType, err string) AppError {
	return AppError{
		Type:   code,
		Reason: err,
	}
}

func (e AppError) Add(err error) AppError {
	fmt.Println(err)
	e.Reason = fmt.Sprintf("%s : %s", e.Reason, err.Error())
	return e
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s, %s", e.Type, e.Reason)
}

var (
	UserNotFoundError        = NewAppError(ErrorNotFound, "user not found")
	CollegeNotFoundError     = NewAppError(ErrorNotFound, "college not found")
	EnvironmentNotFoundError = NewAppError(ErrorNotFound, "environment not found")
	ProblemNotFoundError     = NewAppError(ErrorNotFound, "problem not found")
	SubmissionNotFoundError  = NewAppError(ErrorNotFound, "submission not found")
	SyllabusNotFoundError    = NewAppError(ErrorNotFound, "syllabus not found")
	SubjectNotFoundError     = NewAppError(ErrorNotFound, "subject not found")
	UniversityNotFoundError  = NewAppError(ErrorNotFound, "university not found")
	UserInvalidInput         = NewAppError(ErrorBadData, "invalid input")
	InternalError            = NewAppError(ErrorInternal, "internal server error")
	BadRequest               = NewAppError(ErrorBadData, "bad request")
	NoError                  = NewAppError(ErrorNone, "")
	UnauthorizedError        = NewAppError(ErrorUnauthorized, "unauthorized")
)

type ErrorType string

const (
	ErrorNone          ErrorType = ""
	ErrorTimeout       ErrorType = "timeout"
	ErrorCanceled      ErrorType = "canceled"
	ErrorExec          ErrorType = "execution"
	ErrorBadData       ErrorType = "bad_data"
	ErrorInternal      ErrorType = "internal"
	ErrorUnavailable   ErrorType = "unavailable"
	ErrorNotFound      ErrorType = "not_found"
	ErrorNotAcceptable ErrorType = "not_acceptable"
	ErrorUnauthorized  ErrorType = "unauthorized"
)

type FileType string

const (
	PROBLEM     FileType = "problem"
	TESTFILE    FileType = "test-file"
	SOLUTION    FileType = "solution"
	INITCODE    FileType = "init-code"
	ENVIRONMENT FileType = "environment"
)

type ProgrammingLanguageEnum string

const (
	C          ProgrammingLanguageEnum = "C"
	CPlusPlus  ProgrammingLanguageEnum = "C++"
	Java       ProgrammingLanguageEnum = "Java"
	Python     ProgrammingLanguageEnum = "Python"
	JavaScript ProgrammingLanguageEnum = "JavaScript"
	Go         ProgrammingLanguageEnum = "Go"
	Rust       ProgrammingLanguageEnum = "Rust"
	Text       ProgrammingLanguageEnum = "Text"
	YAML       ProgrammingLanguageEnum = "YAML"
	MYSQL      ProgrammingLanguageEnum = "MYSQL"
	// Add more programming languages as needed
)

func (lang ProgrammingLanguageEnum) GetExtension() string {
	switch lang {
	case C:
		return ".c"
	case CPlusPlus:
		return ".cpp"
	case Java:
		return ".java"
	case Python:
		return ".py"
	case JavaScript:
		return ".js"
	case Go:
		return ".go"
	case Rust:
		return ".rs"
	case Text:
		return ".txt"
	case YAML:
		return ".yml"
	default:
		// Handle unknown language or return a default extension
		return ".txt"
	}
}

type DifficultyEnum string

const (
	DifficultyEasy   DifficultyEnum = "EASY"
	DifficultyMedium DifficultyEnum = "MEDIUM"
	DifficultyHard   DifficultyEnum = "HARD"
)

type TestFilesType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	File     []byte                  `json:"file" validate:"required"`
	Title    string                  `json:"title" validate:"requierd"`
	InitCode []byte                  `json:"init_code" validate:"required"`
}

type ProblemEnvironmentType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Id       uuid.UUID               `json:"id" validate:"required"`
}

// TODO : instead of list maybe have map
type EnvironmentJSON []ProblemEnvironmentType

type TestLinkType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Link     string                  `json:"link" validate:"required"`
	Title    string                  `json:"title" validate:"required"`
}

// TODO : instead of list maybe have map
type TestLinkJSON []TestLinkType

// Problem struct
type Problem struct {
	ID          uuid.UUID       `json:"id" validate:"required"`
	Title       string          `json:"title" validate:"required"`
	CreatedBy   uuid.UUID       `json:"created_by" validate:"required"`
	CreatedAt   time.Time       `json:"created_at" validate:"required"`
	Environment EnvironmentJSON `json:"environment" validate:"required"`
	ProblemLink string          `json:"problem_link" validate:"required"`
	Difficulty  DifficultyEnum  `json:"difficulty" validate:"required"`
	SyllabusID  uuid.UUID       `json:"syllabus_id" validate:"required"`
	TestLinks   TestLinkJSON    `json:"test_links" validate:"required"`
	ProblemFile []byte          `json:"problem_file" validate:"required" gorm:"-"`
	TestFiles   []TestFilesType `json:"test_files" validate:"required" gorm:"-"`
}

// CreateProblemAPIRequest struct
type CreateProblemAPIRequest struct {
	Title       string          `json:"title" validate:"required"`
	CreatedBy   uuid.UUID       `json:"created_by" validate:"required"`
	Environment EnvironmentJSON `json:"environment" validate:"required"`
	ProblemFile []byte          `json:"problem_file" validate:"required"`
	Difficulty  DifficultyEnum  `json:"difficulty" validate:"required"`
	SyllabusID  uuid.UUID       `json:"syllabus_id" validate:"required"`
	TestFiles   []TestFilesType `json:"test_files" validate:"required"`
}

// Value implements the driver.Valuer interface
func (e EnvironmentJSON) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// Scan implements the sql.Scanner interface
func (e *EnvironmentJSON) Scan(value interface{}) error {
	if value == nil {
		*e = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var env []ProblemEnvironmentType
		if err := json.Unmarshal(v, &env); err != nil {
			return err
		}
		*e = EnvironmentJSON(env)
		return nil
	default:
		return errors.New("unsupported type for EnvironmentJSON")
	}
}

// Value implements the driver.Valuer interface
func (tl TestLinkJSON) Value() (driver.Value, error) {
	return json.Marshal(tl)
}

func (e *ProblemEnvironmentType) Scan(value interface{}) error {
	if value == nil {
		*e = ProblemEnvironmentType{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var env []ProblemEnvironmentType
		if err := json.Unmarshal(v, &env); err != nil {
			return err
		}
		// Todo: code might break in case of empty
		*e = env[0]
		return nil
	default:
		return errors.New("unsupported type for ProblemEnvironmentType")
	}
}

// Value implements the driver.Valuer interface
func (tl ProblemEnvironmentType) Value() (driver.Value, error) {
	return json.Marshal(tl)
}

// Scan implements the sql.Scanner interface
func (tl *TestLinkJSON) Scan(value interface{}) error {
	if value == nil {
		*tl = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var testLinks []TestLinkType
		if err := json.Unmarshal(v, &testLinks); err != nil {
			return err
		}
		*tl = TestLinkJSON(testLinks)
		return nil
	default:
		return errors.New("unsupported type for TestLinkJSON")
	}
}

// ProblemAPIResponse struct
type ProblemAPIResponse struct {
	Message *Problem
}

// Implement the Parse method for POST request for CreateProblemAPIRequest
func (r *CreateProblemAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for ProblemAPIResponse
func (pr *ProblemAPIResponse) Write(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(pr)
}

type ListProblemsAPIResponse struct {
	Message []*Problem
}

// Implement the Write method for ListProblemsAPIResponse
func (pr *ListProblemsAPIResponse) Write(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(pr)
}

func NewListProblemsAPIResponse(problems []*Problem) *ListProblemsAPIResponse {
	return &ListProblemsAPIResponse{
		Message: problems,
	}
}

// todo : rename
type InitProblemCode struct {
	Message []byte
}

// Implement the Write method for ListProblemsAPIResponse
func (pr *InitProblemCode) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewInitProblemCode(code []byte) *InitProblemCode {
	return &InitProblemCode{
		Message: code,
	}
}

func (r *CreateProblemAPIRequest) ToProblem() *Problem {
	return &Problem{
		ID:          uuid.New(),
		Title:       r.Title,
		CreatedBy:   r.CreatedBy,
		CreatedAt:   time.Now(),
		Environment: r.Environment,
		ProblemFile: r.ProblemFile,
		Difficulty:  r.Difficulty,
		SyllabusID:  r.SyllabusID,
		TestFiles:   r.TestFiles,
	}
}

func NewCreateProblemAPIResponse(problem *Problem) *ProblemAPIResponse {
	return &ProblemAPIResponse{
		Message: problem,
	}
}

func (e *TestLinkType) Scan(value interface{}) error {
	if value == nil {
		*e = TestLinkType{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var t []TestLinkType
		if err := json.Unmarshal(v, &t); err != nil {
			return err
		}
		// Todo: code might break in case of empty
		*e = t[0]
		return nil
	default:
		return errors.New("unsupported type for TestLinkType")
	}
}

// Value implements the driver.Valuer interface
func (tl TestLinkType) Value() (driver.Value, error) {
	return json.Marshal(tl)
}

// Subject struct
type Subject struct {
	ID           uuid.UUID       `json:"id" validate:"required"`
	Title        string          `json:"title" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
}

// CreateSubjectAPIRequest struct
type CreateSubjectAPIRequest struct {
	Title        string          `json:"title" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
}

// SubjectAPIResponse struct
type SubjectAPIResponse struct {
	Message *Subject
}

// Implement the Parse method for POST request for CreateSubjectAPIRequest
func (r *CreateSubjectAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for SubjectAPIResponse
func (sr *SubjectAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(sr)
}

func (r *CreateSubjectAPIRequest) ToSubject() *Subject {
	return &Subject{
		ID:           uuid.New(),
		Title:        r.Title,
		Description:  r.Description,
		UniversityID: r.UniversityID,
	}
}

func NewCreateSubjectAPIResponse(subject *Subject) *SubjectAPIResponse {
	return &SubjectAPIResponse{
		Message: subject,
	}
}

type ListSubjectsAPIResponse struct {
	Message []*Subject
}

// Implement the Write method for ListSubjectsAPIResponse
func (pr *ListSubjectsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSubjectsAPIResponse(subjects []*Subject) *ListSubjectsAPIResponse {
	return &ListSubjectsAPIResponse{
		Message: subjects,
	}
}

type Status string

const (
	Accepted            Status = "Accepted"              // normal
	MemoryLimitExceeded Status = "Memory Limit Exceeded" // mle
	TimeLimitExceeded   Status = "Time Limit Exceeded"   // tle
	OutputLimitExceeded Status = "Output Limit Exceeded" // ole
	FileError           Status = "File Error"            // fe
	NonzeroExitStatus   Status = "Nonzero Exit Status"
	Signalled           Status = "Signalled"
	InternalErrorStatus Status = "Internal Error" // system error
	Queued              Status = "Queued"
	Running             Status = "Running"
)

// Submission struct
type Submission struct {
	ID        uuid.UUID               `json:"id" validate:"required" gorm:"column:id"`
	ProblemID uuid.UUID               `json:"problem_id" validate:"required" gorm:"column:problem_id;tableName:submissions"`
	Link      string                  `json:"link" validate:"required"`
	CreatedBy uuid.UUID               `json:"created_by" validate:"required"`
	CreatedAt time.Time               `json:"created_at" validate:"required"`
	Score     float64                 `json:"score" validate:"required,min=0,max=100"`
	RunTime   string                  `json:"run_time" validate:"required"`
	Metadata  json.RawMessage         `json:"metadata" validate:"required"`
	Lang      ProgrammingLanguageEnum `json:"lang" validate:"required"`
	Status    Status                  `json:"status"`
	Solution  []byte                  `json:"solution" gorm:"-"`
}

type SubmissionWithProblemTitle struct {
	ID           uuid.UUID               `json:"id"`
	ProblemID    uuid.UUID               `json:"problem_id"`
	Link         string                  `json:"link"`
	CreatedBy    uuid.UUID               `json:"created_by"`
	CreatedAt    time.Time               `json:"created_at"`
	Score        float64                 `json:"score"`
	RunTime      string                  `json:"run_time"`
	Metadata     json.RawMessage         `json:"metadata"`
	Lang         ProgrammingLanguageEnum `json:"lang"`
	Status       Status                  `json:"status"`
	Solution     []byte                  `json:"-"`
	ProblemTitle string                  `json:"problem_title" gorm:"column:problemtitle"`
}

func (s *Submission) UpdateFrom(us Submission) {
	if us.Score > 0 && us.Score < 100 {
		s.Score = us.Score
	}
	if us.RunTime != "" {
		s.RunTime = us.RunTime
	}
	if us.Metadata != nil {
		s.Metadata = us.Metadata
	}
	if us.Status != "" {
		s.Status = us.Status
	}
}

// CreateSubmissionAPIRequest struct
// Todo ; change Lang to Language / add status in sql

type CreateSubmissionAPIRequest struct {
	ProblemID uuid.UUID               `json:"problem_id" validate:"required"`
	Solution  []byte                  `json:"solution" validate:"required"`
	CreatedBy uuid.UUID               `json:"created_by" validate:"required"`
	Metadata  json.RawMessage         `json:"metadata"`
	Lang      ProgrammingLanguageEnum `json:"lang" validate:"required"`
}

// SubmissionAPIResponse struct
type SubmissionAPIResponse struct {
	Message *Submission
}

type UpdateSubmissionAPIRequest struct {
	Score    float64         `json:"score" validate:"min=0,max=100"`
	RunTime  string          `json:"run_time"`
	Metadata json.RawMessage `json:"metadata"`
	Status   Status          `json:"status"`
}

func (r *UpdateSubmissionAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

func (r *UpdateSubmissionAPIRequest) ToSubmissions() *Submission {
	return &Submission{
		Status:   r.Status,
		RunTime:  r.RunTime,
		Metadata: r.Metadata,
		Score:    r.Score,
	}
}

func (r *CreateSubmissionAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for SubmissionAPIResponse
func (sr *SubmissionAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(sr)
}

func (r *CreateSubmissionAPIRequest) ToSubmissions() *Submission {
	return &Submission{
		ID:        uuid.New(),
		ProblemID: r.ProblemID,
		Solution:  r.Solution,
		CreatedBy: r.CreatedBy,
		CreatedAt: time.Now(),
		Metadata:  r.Metadata,
		Lang:      r.Lang,
		Status:    Queued,
	}
}

func NewCreateSubmissionAPIResponse(submission *Submission) *SubmissionAPIResponse {
	return &SubmissionAPIResponse{
		Message: submission,
	}
}

func NewUpdateSubmissionAPIResponse(submission *Submission) *SubmissionAPIResponse {
	return &SubmissionAPIResponse{
		Message: submission,
	}
}

type ListSubmissionsAPIResponse struct {
	Message []*Submission
}

// Implement the Write method for ListSubmissionsAPIResponse
func (pr *ListSubmissionsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSubmissionsAPIResponse(submissions []*Submission) *ListSubmissionsAPIResponse {
	return &ListSubmissionsAPIResponse{
		Message: submissions,
	}
}

type ListSubmissionsWithProbTitleAPIResponse struct {
	Message []*SubmissionWithProblemTitle
}

// Implement the Write method for ListSubmissionsWithProbTitleAPIResponse
func (pr *ListSubmissionsWithProbTitleAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSubmissionsWithProbTitleAPIResponse(submissions []*SubmissionWithProblemTitle) *ListSubmissionsWithProbTitleAPIResponse {
	return &ListSubmissionsWithProbTitleAPIResponse{
		Message: submissions,
	}
}

type SyllabusLevelEnum string

const (
	SyllabusLevelUniversity SyllabusLevelEnum = "UNIVERSITY"
	SyllabusLevelCollege    SyllabusLevelEnum = "COLLEGE"
	SyllabusLevelGlobal     SyllabusLevelEnum = "GLOBAL"
)

// AccessLevelModeEnum represents the lab_rank.access_level_mode_enum enum type.
// Syllabus struct
type Syllabus struct {
	ID            uuid.UUID         `json:"id" validate:"required"`
	SubjectID     uuid.UUID         `json:"subject_id" validate:"required"`
	UniCollegeID  uuid.UUID         `json:"uni_college_id" validate:"required"`
	SyllabusLevel SyllabusLevelEnum `json:"syllabus_level" validate:"required"`
}

// CreateSyllabusAPIRequest struct
type CreateSyllabusAPIRequest struct {
	SubjectID     uuid.UUID         `json:"subject_id" validate:"required"`
	UniCollegeID  uuid.UUID         `json:"uni_college_id" validate:"required"`
	SyllabusLevel SyllabusLevelEnum `json:"syllabus_level" validate:"required"`
}

// SyllabusAPIResponse struct
type SyllabusAPIResponse struct {
	Message *Syllabus
}

// Implement the Parse method for POST request for CreateSyllabusAPIRequest
func (r *CreateSyllabusAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for SyllabusAPIResponse
func (sr *SyllabusAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(sr)
}

func (r *CreateSyllabusAPIRequest) ToSyllabus() *Syllabus {
	return &Syllabus{
		ID:            uuid.New(),
		SubjectID:     r.SubjectID,
		UniCollegeID:  r.UniCollegeID,
		SyllabusLevel: r.SyllabusLevel,
	}
}

func NewCreateSyllabusAPIResponse(syllabus *Syllabus) *SyllabusAPIResponse {
	return &SyllabusAPIResponse{
		Message: syllabus,
	}
}

type ListSyllabusAPIResponse struct {
	Message []*Syllabus
}

// Implement the Write method for ListSyllabussAPIResponse
func (pr *ListSyllabusAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSyllabusAPIResponse(syllabus []*Syllabus) *ListSyllabusAPIResponse {
	return &ListSyllabusAPIResponse{
		Message: syllabus,
	}
}

func (r Subject) ToSyllabus(id uuid.UUID, level SyllabusLevelEnum) *Syllabus {
	return &Syllabus{
		ID:            uuid.New(),
		SubjectID:     r.ID,
		UniCollegeID:  id,
		SyllabusLevel: level,
	}
}

// University struct
type University struct {
	ID          uuid.UUID       `json:"id" validate:"required"`
	Title       string          `json:"title" validate:"required"`
	Description json.RawMessage `json:"description" validate:"required"`
}

// CreateUniversityAPIRequest struct
type CreateUniversityAPIRequest struct {
	Title       string          `json:"title" validate:"required"`
	Description json.RawMessage `json:"description" validate:"required"`
}

// UniversityAPIResponse struct
type UniversityAPIResponse struct {
	Message *University
}

// Implement the Parse method for POST request for CreateUniversityAPIRequest
func (r *CreateUniversityAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for UniversityAPIResponse
func (ur *UniversityAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(ur)
}

func (r *CreateUniversityAPIRequest) ToUniversity() *University {
	return &University{
		ID:          uuid.New(),
		Title:       r.Title,
		Description: r.Description,
	}
}

func NewCreateUniversityAPIResponse(university *University) *UniversityAPIResponse {
	return &UniversityAPIResponse{
		Message: university,
	}
}

type UniversityIdName struct {
	ID    uuid.UUID
	Title string
}

func NewUniversityIdName(id uuid.UUID, name string) *UniversityIdName {
	return &UniversityIdName{
		ID:    id,
		Title: name,
	}
}

type ListUniversitiesIdNamesAPIResponse struct {
	Message []*UniversityIdName
}

type ListUniversitiesAPIResponse struct {
	Message []*University
}

// Implement the Write method for ListUniversitiesAPIResponse
func (pr *ListUniversitiesIdNamesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

// Implement the Write method for ListUniversitiesAPIResponse
func (pr *ListUniversitiesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}
func NewListUniversitiesIdNamesAPIResponse(universities []*UniversityIdName) *ListUniversitiesIdNamesAPIResponse {
	if universities == nil {
		universities = []*UniversityIdName{} // Initialize an empty slice if it's nil
	}
	return &ListUniversitiesIdNamesAPIResponse{
		Message: universities,
	}
}

func NewListUniversitiesAPIResponse(universities []*University) *ListUniversitiesAPIResponse {
	return &ListUniversitiesAPIResponse{
		Message: universities,
	}
}
