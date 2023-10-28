package user

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)


type UserService interface {
	Create(context.Context,  *models.User) (*models.User, error)
	// Update(models.ServiceRequest) models.ServiceResponse
	// Fetch(models.ServiceRequest) models.ServiceResponse
	// Delete(models.ServiceRequest) models.ServiceResponse
}


type userService struct {
	// You can add any dependencies or data storage components here
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}