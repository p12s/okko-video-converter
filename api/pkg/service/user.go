package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
)

const tokenTTL = 24 * time.Hour * 365

type tokenClaims struct {
	jwt.StandardClaims
	Code uuid.UUID `json:"code"`
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser() (uuid.UUID, error) {
	userCode := uuid.New()
	err := s.repo.CreateUser(userCode)
	if err != nil {
		return uuid.Nil, err
	}

	return userCode, nil
}

func (s *UserService) GenerateToken(userCode uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userCode,
	})
	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (s *UserService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Code, nil
}
