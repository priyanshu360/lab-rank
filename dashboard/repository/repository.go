package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type UserRepository interface {
	CreateUser(context.Context, models.User)  error
    GetUserByID(context.Context, uuid.UUID) (models.User, error)
    GetUserByEmail(context.Context, string) (models.User, error)
    UpdateUser(context.Context, uuid.UUID, models.User) error
    DeleteUser(context.Context, uuid.UUID) error
    ListUsers(context.Context, int, int) ([]models.User, error)
}
