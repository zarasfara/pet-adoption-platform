package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordMismatch    = errors.New("passwords do not match")
	ErrInvalidTokenMethod  = errors.New("invalid token method")
	ErrUserNotFound        = errors.New("user not found")
)

type AuthService struct {
	repo           repository.Authorization
	accessTokenTTL time.Duration
	signingToken   []byte
}

func NewAuthService(repo repository.Authorization, cfg *config.Config) *AuthService {
	return &AuthService{
		repo:           repo,
		accessTokenTTL: cfg.JWT.AccessTokenTTL,
		signingToken:   []byte(cfg.JWT.SigningToken),
	}
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s AuthService) CreateUser(user models.AddRecordUser) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if !checkPasswordHash(password, user.Password) {
		return "", ErrPasswordMismatch
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(s.accessTokenTTL).Unix(),
		"iat": time.Now().Unix(),
		"sub": user.Id,
	})

	signedToken, err := token.SignedString(s.signingToken)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenMethod
		}
		return s.signingToken, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrTokenInvalidClaims
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return 0, jwt.ErrTokenInvalidSubject
	}

	return int(userId), nil
}

func (s AuthService) GetCurrentUser(userID int) (models.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return models.User{}, ErrUserNotFound
	}

	return user, nil
}
