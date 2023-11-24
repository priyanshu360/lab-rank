package user

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type UserService interface {
	Create(context.Context, *models.User) (*models.User, models.AppError)
	Update(context.Context, *models.UpdateUserAPIRequest) (*models.User, models.AppError)
	Fetch(context.Context, *models.GetUserAPIRequest) ([]*models.User, models.AppError)
	Delete(context.Context, string) models.AppError
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

func (s *userService) Fetch(ctx context.Context, req *models.GetUserAPIRequest) ([]*models.User, models.AppError) {
	var users []*models.User
	switch {
	case req.EmailID != "":
		emailID := req.EmailID
		if user, err := s.repo.GetUserByEmail(ctx, emailID); err != models.NoError {
			return users, err
		} else {
			users = append(users, &user)
			return users, models.NoError
		}

	case req.UserID != "":
		if userID, err := uuid.Parse(req.UserID); err != nil {
			return nil, models.InternalError.Add(err)
		} else {
			if user, err := s.repo.GetUserByID(ctx, userID); err != models.NoError {
				return users, err
			} else {
				users = append(users, &user)
				return users, models.NoError
			}
		}
	case req.Limit != "":
		if limit, err := strconv.ParseInt(req.Limit, 10, 64); err != nil {
			return s.repo.ListUsersWisthLimit(ctx, 1, 10)
		} else {
			return s.repo.ListUsersWisthLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.ListUsersWisthLimit(ctx, 1, 10)

	}
}

func (s *userService) Update(ctx context.Context, request *models.UpdateUserAPIRequest) (*models.User, models.AppError) {

	defaultUser, err := s.repo.GetUserByID(ctx, request.ID)
	if err != models.NoError {
		return nil, err
	}
	updatedUser := request.ToUser(defaultUser)
	if err := s.repo.UpdateUser(ctx, request.ID, *updatedUser); err != models.NoError {
		return nil, err
	}

	return updatedUser, models.NoError
}

func (s *userService) Delete(ctx context.Context, id string) models.AppError {
	if userID, err := uuid.Parse(id); err != nil {
		return models.InternalError.Add(err)
	} else {
		if err := s.repo.DeleteUser(ctx, userID); err != models.NoError {
			return err
		}

		return models.NoError
	}

}
