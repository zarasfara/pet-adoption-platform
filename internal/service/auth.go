package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo            repository.Authorization
	refreshTokenTTL time.Duration
	signingToken    string
}

func NewAuthService(repo repository.Authorization, cfg *config.Config) *AuthService {
	return &AuthService{
		repo:            repo,
		refreshTokenTTL: cfg.JWT.RefreshTokenTTL,
		signingToken:    cfg.JWT.SigningToken,
	}
}

func hashPassword(password string) string {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s AuthService) CreateUser(user models.User) error {
	user.Password = hashPassword(user.Password)

	return s.repo.CreateUser(user)
}

func (s AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, hashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"exp": time.Now().Add(s.refreshTokenTTL).Unix(),
		"iat": time.Now().Unix(),
		"sub": user.Id,
	})

	signedToken, err := token.SignedString([]byte(s.signingToken))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
