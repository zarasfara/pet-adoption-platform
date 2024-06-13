package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"github.com/zarasfara/pet-adoption-platform/internal/repository"
	"github.com/zarasfara/pet-adoption-platform/pkg/auth"
	"github.com/google/uuid"
)

var (
	ErrPasswordMismatch   = errors.New("passwords do not match")
	ErrInvalidTokenMethod = errors.New("invalid token method")
	ErrUserNotFound       = errors.New("user not found")
)

type Authorization interface {
	CreateUser(user models.AddRecordUser) error
	SignIn(email, password string) (string, string, error)
	UserByID(userId int) (models.User, error)
	RegenerateTokens(token string) (accessToken string, refreshToken string, err error)
	UserIDFromToken(accessToken string) (int, error)
}

type authService struct {
	repo            repository.Authorization
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	signingToken    []byte
}

func NewAuthService(repo repository.Authorization, cfg *config.Config) *authService {
	return &authService{
		repo:            repo,
		accessTokenTTL:  cfg.JWT.AccessTokenTTL,
		refreshTokenTTL: cfg.JWT.RefreshTokenTTL,
		signingToken:    []byte(cfg.JWT.SigningToken),
	}
}

func (s authService) CreateUser(user models.AddRecordUser) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s authService) generateAccessToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(userId),
	})

	signedToken, err := token.SignedString(s.signingToken)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s authService) generateRefreshToken(userId int) (string, error) {
	session := models.RefreshSession{
		UserId: userId,
		RefreshToken: uuid.NewString(),
		ExpiresIn:    time.Now().Add(s.refreshTokenTTL).Unix(),
		CreatedAt:    time.Now(),
	}

	err := s.repo.CreateRefreshSession(session)
	if err != nil {
		return "", err
	}

	return session.RefreshToken, nil
}

func (s authService) SignIn(email, password string) (string, string, error) {
	user, err := s.repo.UserByEmail(email)
	if err != nil {
		return "", "", err
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", "", ErrPasswordMismatch
	}

	accessToken, err := s.generateAccessToken(user.Id)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateRefreshToken(user.Id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s authService) GenerateTokens(userId int) (accessToken, refreshToken string, err error) {
	accessToken, err = s.generateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.generateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s authService) parseAccessToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidTokenMethod
		}
		return s.signingToken, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}

func (s authService) UserIDFromToken(accessToken string) (int, error) {
	claims, err := s.parseAccessToken(accessToken)
	if err != nil {
		return 0, err
	}

	registeredClaims, ok := claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	userID, err := strconv.Atoi(registeredClaims.Subject)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s authService) RegenerateTokens(refreshToken string) (access string, refresh string, err error) {
	// получаем запись рефреш-сессии по UUID'у рефреш токена
	refreshSession, err := s.repo.GetRefreshSessionByUUID(refreshToken)
	if err != nil {
		return "", "", err
	}

	// Сохраняем текущую рефреш-сессию в переменную и удаляет ее из таблицы
	err = s.repo.DeleteRefreshSessionByUUID(refreshSession.RefreshToken)
	if err != nil {
		return "", "", nil
	}

	if time.Now().Unix() > refreshSession.ExpiresIn {
		return "", "", errors.New("refresh token has expired")
	}

	access, refresh, err = s.GenerateTokens(refreshSession.UserId)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s authService) UserByID(userID int) (models.User, error) {
	user, err := s.repo.UserByID(userID)
	if err != nil {
		return models.User{}, ErrUserNotFound
	}

	return user, nil
}
