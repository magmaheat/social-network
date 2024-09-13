package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"sn-auth/internal/entity"
	"sn-auth/internal/repo"
	"sn-auth/internal/repo/repoerrs"
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
	user := entity.User{
		Username: input.Username,
		Password: a.passwordHasher.Hash(input.Password),
	}

	userId, err := a.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrUserAlreadyExists
		}
		log.Errorf("AuthService.CreateUser - a.userRepo.CreateUser: %v", err)
	}
	return userId, nil
}

func (a *AuthService) GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error) {
	user, err := a.userRepo.GetUserByUsernameAndPassword(ctx, input.Username, a.passwordHasher.Hash(input.Password))
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrUserNotFound
		}
		log.Errorf("AuthService.GenerateToken: cannot get user: %v", err)
		return "", ErrCannotGetUser
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	tokenString, err := token.SignedString([]byte(a.signKey))
	if err != nil {
		log.Errorf("AuthService.GenerateToken: cannot sing token: %s", err)
		return "", ErrCannotSignToken
	}

	return tokenString, nil

}

func (a *AuthService) ParseToken(token string) (int, error) { return 0, nil }
