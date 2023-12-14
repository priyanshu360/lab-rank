package postgres

import (
	"context"
	"errors"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type authPostgres struct {
	db *gorm.DB
}

// NewauthPostgresRepo creates a new PostgreSQL repository for auths.
func NewAuthPostgresRepo(db *gorm.DB) *authPostgres {
	return &authPostgres{db}
}

func (r *authPostgres) SignUp(ctx context.Context, user models.User, auth models.Auth) models.AppError {
	// Check if the email is already registered
	var existingUser models.User
	result := r.db.Where("email = ?", user.Email).Table("lab_rank.user").First(&existingUser)
	if result.Error == nil {
		return models.InternalError.Add(errors.New("user already exist"))
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		return models.InternalError.Add(tx.Error)
	}

	// Insert user data into the "user" table
	if err := tx.Table("lab_rank.user").Create(&user).Error; err != nil {
		tx.Rollback()
		return models.InternalError.Add(err)
	}

	// Insert authentication data into the "auth" table
	if err := tx.Table("lab_rank.auth").Create(&auth).Error; err != nil {
		tx.Rollback()
		return models.InternalError.Add(err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		models.InternalError.Add(err)
	}

	return models.NoError
}

func (psql *authPostgres) GetUserAuthByEmail(ctx context.Context, email string) (*models.User, *models.Auth, models.AppError) {
	var user models.User
	var auth models.Auth
	result := psql.db.WithContext(ctx).Where("email = ?", email).Table("lab_rank.user").First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return &user, &auth, models.UserNotFoundError
		}
		return &user, &auth, models.InternalError.Add(result.Error)
	}

	result = psql.db.WithContext(ctx).Where("user_id = ?", user.ID).Table("lab_rank.auth").First(&auth)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found
			return &user, &auth, models.UserNotFoundError.Add(result.Error)
		}
		return &user, &auth, models.InternalError.Add(result.Error)
	}

	return &user, &auth, models.NoError
}
