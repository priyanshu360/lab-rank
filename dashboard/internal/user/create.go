package user

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

func (s *userService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println(user)

	user.ID = uuid.New()
	user.Status = models.UserStatusInactive

	if err := user.GenerateReqFields(); err != nil {
		log.Println("error 1", err)
		return nil, err
	}

	if err := s.repo.CreateUser(ctx, *user); err != nil {
		log.Println("error 2", err)

		return nil, err
	}
    
	log.Println(user)
	return user, nil
}