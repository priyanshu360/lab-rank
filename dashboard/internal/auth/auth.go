package auth

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(context.Context, string, string) (*models.LoginAPIResponse, models.AppError)
	// Logout(context.Context, string) error
	SignUp(context.Context, *models.User, string) models.AppError
}

type service struct {
	// You can add any dependencies or data storage components here
	syllabus syllabus.Service
	repo     repository.AuthRepository
}

func New(repo repository.AuthRepository, syllabus syllabus.Service) *service {
	return &service{
		repo:     repo,
		syllabus: syllabus,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (*models.LoginAPIResponse, models.AppError) {
	// Authenticate user (validate email and password)
	user, auth, err := s.repo.GetUserAuthByEmail(ctx, email)
	if err != models.NoError {
		return nil, models.InternalError.Add(errors.New("Login failed"))
	}

	// Validate password
	if !verifyPassword(password, auth.PasswordHash, string(auth.Salt)) {
		return nil, models.InternalError.Add(errors.New("Login failed"))
	}

	// Create JWT token
	token, jwtErr := generateJWTToken(user, auth)
	if jwtErr != nil {
		return nil, models.InternalError.Add(jwtErr)
	}

	return models.NewLoginAPIResponse(token), models.NoError
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
		AccessIDs:    []uuid.UUID{},
		Salt:         salt,
		PasswordHash: passwordHash,
	}

	if err := s.repo.SignUp(ctx, *user, auth); err != models.NoError {
		return err
	}

	s.syllabus.UpdateAccessIDsForUser(context.Background(), user)
	// Call the repository to store the user and authentication entry
	return models.NoError
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

func generateJWTToken(user *models.User, auth *models.Auth) (string, error) {
	// Replace the following with your own secret key and token expiration time
	secretKey := []byte("your_secret_key")
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	accesses, err := json.Marshal(auth.AccessIDs)
	if err != nil {
		return "", err
	}

	claims := &jwt.RegisteredClaims{
		ID:        user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   string(accesses),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
