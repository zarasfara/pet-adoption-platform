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

func (s *AuthService) CreateUser(user models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return "", err
	}
	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("неверный пароль")
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

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("ошибка: неверный метод токена")
		}
		return s.signingToken, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("ошибка: невалидный токен")
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("ошибка: неверный формат идентификатора пользователя")
	}

	return int(userId), nil
}
