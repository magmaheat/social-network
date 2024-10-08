package service

import (
	"context"
	"github.com/magmaheat/social-network/sn-auth/internal/repo"
	"github.com/magmaheat/social-network/sn-auth/pkg/hasher"
	"time"
)

type AuthCreateUserInput struct {
	Username string
	Password string
}

type AuthGenerateTokenInput struct {
	Username string
	Password string
}

type Auth interface {
	CreateUser(ctx context.Context, input AuthCreateUserInput) (int, error)
	GenerateToken(ctx context.Context, input AuthGenerateTokenInput) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Auth
}

type ServicesDependencies struct {
	Repos  *repo.Repositories
	Hasher hasher.PasswordHasher

	SignKey  string
	TokenTTL time.Duration
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Auth: NewAuthService(deps.Repos, deps.Hasher, deps.SignKey, deps.TokenTTL),
	}
}
