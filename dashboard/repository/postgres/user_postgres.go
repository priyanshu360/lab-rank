package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type userPostgres struct {
	db *gorm.DB
}

// NewUserPostgresRepo creates a new PostgreSQL repository for users.
func NewUserPostgresRepo(db *gorm.DB) *userPostgres {
	return &userPostgres{db}
}

// GetUserByID retrieves a user by their user ID.
func (psql *userPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, models.AppError) {
	var user models.User
	result := psql.db.WithContext(ctx).Table("lab_rank.user").First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return user, models.UserNotFoundError
		}
		return user, models.InternalError.Add(result.Error)
	}
	return user, models.NoError
}

// GetUserByEmail retrieves a user by their email.
func (psql *userPostgres) GetUserByEmail(ctx context.Context, email string) (models.User, models.AppError) {
	var user models.User
	result := psql.db.WithContext(ctx).Where("email = ?", email).Table("lab_rank.user").First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return user, models.UserNotFoundError
		}
		return user, models.InternalError.Add(result.Error)
	}
	return user, models.NoError
}

// UpdateUser updates a user's information.
func (psql *userPostgres) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) models.AppError {
	// Check if the user with the provided ID exists before updating
	var existingUser models.User
	result := psql.db.WithContext(ctx).First(&existingUser, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return models.UserNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the update
	result = psql.db.WithContext(ctx).Model(&user).Where("id = ?", userID).Updates(user)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// DeleteUser deletes a user by their user ID.
func (psql *userPostgres) DeleteUser(ctx context.Context, userID uuid.UUID) models.AppError {
	// Check if the user with the provided ID exists before deletion
	var existingUser models.User
	result := psql.db.WithContext(ctx).First(&existingUser, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return models.UserNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the deletion
	result = psql.db.WithContext(ctx).Delete(&models.User{}, userID)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

func (psql *userPostgres) ListUsers(ctx context.Context, page int, pageSize int) ([]models.User, models.AppError) {
	var users []models.User
	result := psql.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}
	return users, models.NoError
}

func (psql *userPostgres) CreateUser(ctx context.Context, user models.User) models.AppError {
	result := psql.db.WithContext(ctx).Table("lab_rank.user").Create(user)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	log.Println(user, result)
	return models.NoError
}
