package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(context.Context, string, string) (string, models.AppError)
	// Logout(context.Context, string) error
	SignUp(context.Context, *models.User, string) models.AppError
}

type service struct {
	// You can add any dependencies or data storage components here
	repo repository.AuthRepository
}

func New(repo repository.AuthRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (string, models.AppError) {
	// Authenticate user (validate email and password)
	user, auth, err := s.repo.GetUserByEmail(ctx, email)
	if err != models.NoError {
		return "", models.InternalError.Add(errors.New("Login failed"))
	}

	// Validate password
	if !verifyPassword(password, auth.PasswordHash, string(auth.Salt)) {
		return "", models.InternalError.Add(errors.New("Login failed"))
	}

	// Create JWT token
	token, jwtErr := generateJWTToken(user)
	if jwtErr != nil {
		return "", models.InternalError.Add(jwtErr)
	}

	return token, models.NoError
}

func (s *service) SignUp(ctx context.Context, user *models.User, password string) models.AppError {
	// Implement any business logic for signup (e.g., validation)

	user.ID = uuid.New()
	user.Status = models.UserStatusInactive

	if err := user.GenerateReqFields(); err != nil {
		return models.InternalError.Add(err)
	}

	// Generate salt
	salt, err := generateSalt()
	if err != nil {
		return models.InternalError.Add(err)
	}

	// Hash the password
	passwordHash, err := hashPassword(password, string(salt))
	if err != nil {
		return models.InternalError.Add(err)
	}

	auth := models.Auth{
		UserID:       user.ID,
		AccessIDs:    "{}",
		Salt:         salt,
		PasswordHash: passwordHash,
	}

	// Call the repository to store the user and authentication entry
	return s.repo.SignUp(ctx, *user, auth)
}

// generateSalt generates a random salt using the bcrypt library.
func generateSalt() ([]byte, error) {
	saltBytes := make([]byte, 32) // Adjust the size of the salt as needed
	_, err := rand.Read(saltBytes)
	if err != nil {
		return nil, err
	}

	return saltBytes, nil
}

// hashPassword hashes the given password using the bcrypt library and the provided salt.
func hashPassword(password, salt string) (string, error) {
	passwordBytes := []byte(password + salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func verifyPassword(password, hashedPassword, salt string) bool {
	passwordBytes := []byte(password + salt)
	hashedPasswordBytes := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	return err == nil
}

func generateJWTToken(user *models.User) (string, error) {
	// Replace the following with your own secret key and token expiration time
	secretKey := []byte("your_secret_key")
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	claims := &jwt.StandardClaims{
		Subject:   user.ID.String(),
		ExpiresAt: expirationTime.Unix(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
