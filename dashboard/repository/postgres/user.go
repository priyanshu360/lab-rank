package postgres

import (
	"context"

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
func (psql *userPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
    var user models.User
    result := psql.db.WithContext(ctx).First(&user, userID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // User not found
            return user, models.ErrUserNotFound
        }
        return user, result.Error
    }
    return user, nil
}


// GetUserByEmail retrieves a user by their email.
func (psql *userPostgres) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
    var user models.User
    result := psql.db.WithContext(ctx).Where("email = ?", email).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // User not found
            return user, models.ErrUserNotFound
        }
        return user, result.Error
    }
    return user, nil
}

// UpdateUser updates a user's information.
func (psql *userPostgres) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) error {
    // Check if the user with the provided ID exists before updating
    var existingUser models.User
    result := psql.db.WithContext(ctx).First(&existingUser, userID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // User not found
            return models.ErrUserNotFound
        }
        return result.Error
    }

    // Perform the update
    result = psql.db.WithContext(ctx).Model(&user).Where("id = ?", userID).Updates(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


// DeleteUser deletes a user by their user ID.
func (psql *userPostgres) DeleteUser(ctx context.Context, userID uuid.UUID) error {
    // Check if the user with the provided ID exists before deletion
    var existingUser models.User
    result := psql.db.WithContext(ctx).First(&existingUser, userID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // User not found
            return models.ErrUserNotFound
        }
        return result.Error
    }

    // Perform the deletion
    result = psql.db.WithContext(ctx).Delete(&models.User{}, userID)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


func (psql *userPostgres) ListUsers(ctx context.Context, page int, pageSize int) ([]models.User, error) {
    var users []models.User
    result := psql.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (psql *userPostgres) CreateUser(ctx context.Context, user models.User) error {
    result := psql.db.WithContext(ctx).Table("lab_rank.user").Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}