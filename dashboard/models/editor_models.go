package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type EditorStatus int

const (
	Scaled EditorStatus = iota
	Stoped
	Deleted
	Pending
)

type Editor struct {
	ID          int                    `json:"id" bson:"id" auto:"increment"`
	UserID      uuid.UUID              `json:"user_id" bson:"user_id"`
	PodIP       string                 `json:"pod_ip" bson:"pod_ip"`
	Deployment  string                 `json:"deployment" bson:"deployment"`
	Status      EditorStatus           `json:"status" bson:"status"`
	ProblemID   int                    `json:"problem_id" bson:"problem_id"`
	CreatedAt   time.Time              `json:"created_at" bson:"created_at"`
	LastUpdated time.Time              `json:"last_updated" bson:"last_updated"`
	Environment ProblemEnvironmentType `json:"environment" bson:"environment"`
}

type StartEditorReq struct {
	UserID      uuid.UUID              `json:"user_id" bson:"user_id"`
	ProblemID   int                    `json:"problem_id" bson:"problem_id"`
	Environment ProblemEnvironmentType `json:"environment" bson:"environment"`
}

func (req *StartEditorReq) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if req.UserID == uuid.Nil {
		return errors.New("invalid user_id")
	}

	if req.ProblemID == 0 {
		return errors.New("invalid problem_id")
	}

	return nil
}

func (r *StartEditorReq) ToEditor() *Editor {
	now := time.Now()
	return &Editor{
		UserID:      r.UserID,
		ProblemID:   r.ProblemID,
		Environment: r.Environment,
		CreatedAt:   now,
		LastUpdated: now,
		Status:      4, // Default status value
	}
}

type CreateEditorAPIResponse struct {
	Message *Editor
}

func (cr *CreateEditorAPIResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}

func NewCreateEditorAPIResponse(editor *Editor) *CreateEditorAPIResponse {
	return &CreateEditorAPIResponse{
		Message: editor,
	}
}
