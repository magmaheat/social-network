package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/entity"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/repo"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/repo/repoerrs"
	"github.com/magmaheat/social-network/tree/main/sn-auth/pkg/hasher"
	log "github.com/sirupsen/logrus"
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
		return 0, ErrCannotCreateUser
	}

	return userId, nil
}

func (a *AuthService) GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error) {
	user, err := a.userRepo.GetUserByUsernameAndPassword(ctx, input.Username)

	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", ErrUserNotFound
		}

		log.Errorf("AuthService.GenerateToken: cannot get user: %v", err)
		return "", ErrCannotGetUser
	}

	if !a.passwordHasher.CheckPassword(user.Password, input.Password) {
		return "", ErrUserNotFound
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
		log.Errorf("AuthService.GenerateToken: cannot sign token: %v", err)
		return "", ErrCannotSignToken
	}

	return tokenString, nil
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.signKey), nil
	})

	if err != nil {
		return 0, ErrCannotParseToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, ErrCannotParseToken
	}

	return claims.UserId, nil
}
