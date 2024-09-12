package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"sn-auth/internal/repo"
	"sn-auth/pkg/hasher"
	"time"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthService struct {
	userRepo       repo.User
	passwordHasher hasher.PasswordHasher
	signKey        string
	tokenTTL       time.Duration
}

func NewAuthService(userRepo repo.User, passwordHasher hasher.PasswordHasher, signKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}

func (a *AuthService) CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error) {
	return 0, nil
}

func (a *AuthService) GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error) {
	return "", nil
}

func (a *AuthService) ParseToken(token string) (int, error) { return 0, nil }
