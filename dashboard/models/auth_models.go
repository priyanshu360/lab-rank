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
)

type Auth struct {
	UserID       uuid.UUID           `json:"user_id"`
	Salt         []byte              `json:"salt"`
	PasswordHash string              `json:"password_hash"`
	Mode         AccessLevelModeEnum `json:"mode"`
}

type SignUpAPIRequest struct {
	CreateUserAPIRequest
	Password string `json:"password"`
}

type SignUpAPIResponse CreateUserAPIResponse

type LoginAPIResponse struct {
	Jwt           string
	UniversityID string
	CollegeID     uuid.UUID
	UserID       uuid.UUID
}

func NewLoginAPIResponse(jwt, universityID string, userID, collegeID uuid.UUID) *LoginAPIResponse {
	return &LoginAPIResponse{
		Jwt:           jwt,
		UniversityID:  universityID,
		CollegeID:     collegeID,
		UserID:        userID,
	}
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
