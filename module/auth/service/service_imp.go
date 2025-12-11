package service

import (
	"Shop/models"
	"Shop/module/auth/repository"
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) UserService {
	return &service{repo: repo}
}

// Register - Create new user and return JWT token
func (s *service) Register(ctx context.Context, data RegisterUserData) (*models.User, error) {
	// Create user
	user := &models.User{
		Name:  data.Name,
		Email: data.Email,
	}

	// Hash password using bcrypt
	err := user.SetPassword(data.Password)
	if err != nil {
		return nil, err
	}

	// Save to database
	err = s.repo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login - Verify credentials and return JWT token
func (s *service) Login(ctx context.Context, data LoginUserData) (*models.User, error) {
	// Find user by email
	user, err := s.repo.Login(ctx, data.Email, data.Password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if !user.CheckPassword(data.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// GenerateJWT - Generate JWT token for user (helper function)
func GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "A_VERY_STRONG_JWT_SECRET_KEY_12345_SHOP" // Fallback secret
	}
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
