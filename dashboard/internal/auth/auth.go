package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(context.Context, string, string) (*models.LoginAPIResponse, models.AppError)
	Logout(context.Context, string) models.AppError
	Authenticate(context.Context, string) (*models.AuthSession, models.AppError)
	SignUp(context.Context, *models.User, string) models.AppError
}

type service struct {
	// You can add any dependencies or data storage components here
	syllabus syllabus.Service
	repo     repository.AuthRepository
	session  repository.SessionRepository
}

func New(repo repository.AuthRepository, session repository.SessionRepository, syllabus syllabus.Service) *service {
	return &service{
		session:  session,
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

	session := models.NewAuthSession(user, auth.Mode)
	sessionID, err := s.session.SetSession(ctx, session)
	if err != models.NoError {
		return nil, err
	}
	// Create JWT token
	token, jwtErr := generateJWTToken(sessionID)
	if jwtErr != nil {
		return nil, models.InternalError.Add(jwtErr)
	}

	return models.NewLoginAPIResponse(token), models.NoError
}

func (s *service) Authenticate(ctx context.Context, jwt string) (*models.AuthSession, models.AppError) {
	log.Println(jwt)
	sessionID, err := validateJWTToken(jwt)
	if err != nil {
		return nil, models.InternalError.Add(err)
	}

	return s.session.GetSession(ctx, sessionID)
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
		Salt:         salt,
		PasswordHash: passwordHash,
		Mode:         models.AccessLevelStudent,
	}

	if err := s.repo.SignUp(ctx, *user, auth); err != models.NoError {
		return err
	}

	// s.syllabus.UpdateAccessIDsForUser(context.Background(), user)
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

func validateJWTToken(tokenString string) (uuid.UUID, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetJWTKey()), nil
	})

	log.Println("error, ", err)

	if err != nil {
		return uuid.Nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Validate the session ID
		sessionID, err := uuid.Parse(claims.Subject)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid session ID format")
		}

		return sessionID, nil
	}

	return uuid.Nil, fmt.Errorf("invalid token")
}

func generateJWTToken(sessionID uuid.UUID) (string, error) {
	// Replace the following with your own secret key and token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   sessionID.String(),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.GetJWTKey()))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *service) Logout(ctx context.Context, jwtToken string) models.AppError {
	sessionID, err := validateJWTToken(jwtToken)
	if err != nil {
		return models.InternalError.Add(err)
	}

	// Remove the session from the data store (e.g., Redis)
	appErr := s.session.RemoveSession(ctx, sessionID)
	if err != models.NoError {
		return appErr
	}

	return models.NoError
}
