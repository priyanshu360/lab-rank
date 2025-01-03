package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"

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
	UserID       uuid.UUID           `json:"user_id" gorm:"primaryKey"`
	Salt         []byte              `json:"salt" gorm:"type:bytea"`
	PasswordHash string              `json:"password_hash" gorm:"type:string"`
	Mode         AccessLevelModeEnum `json:"mode" gorm:"type:string"`
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
		// DOB:          r.DOB,
		Name:     r.Name,
		UserName: r.UserName,
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
