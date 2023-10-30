package user

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)


type UserService interface {
	Create(context.Context,  *models.User) (*models.User, models.AppError)
	// Update(models.ServiceRequest) models.ServiceResponse
	Fetch(context.Context,  models.GetUserAPIRequest) (*models.User, models.AppError)
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


func (s *userService) Create(ctx context.Context, user *models.User) (*models.User, models.AppError) {

	user.ID = uuid.New()
	user.Status = models.UserStatusInactive

	if err := user.GenerateReqFields(); err != nil {
		log.Println("error 1", err)
		return nil, models.InternalError.Add(err)
	}

	if err := s.repo.CreateUser(ctx, *user); err != models.NoError {
		return nil, err
	}
    
	return user, models.NoError  
}

func (s *userService) Fetch(ctx context.Context, req models.GetUserAPIRequest) (*models.User, models.AppError) {

    switch {
	case req.UserID != uuid.Nil:
        userID := req.UserID
		if user, err := s.repo.GetUserByID(ctx, userID); err != models.NoError {
			return nil, err
		}else {return &user, models.NoError}
			
    case req.EmailID != "":
		emailID := req.EmailID
		if user, err := s.repo.GetUserByEmail(ctx, emailID); err != models.NoError {
			return nil, err
		}else {return &user, models.NoError}
	
	default:
		return nil,models.BadRequest

}	
}