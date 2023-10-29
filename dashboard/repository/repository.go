package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type UserRepository interface {
	CreateUser(context.Context, models.User)  models.AppError
    GetUserByID(context.Context, uuid.UUID) (models.User, models.AppError)
    GetUserByEmail(context.Context, string) (models.User, models.AppError)
    UpdateUser(context.Context, uuid.UUID, models.User) models.AppError
    DeleteUser(context.Context, uuid.UUID) models.AppError
    ListUsers(context.Context, int, int) ([]models.User, models.AppError)
}
