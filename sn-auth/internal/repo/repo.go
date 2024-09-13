package repo

import (
	"context"
	"sn-auth/internal/entity"
	"sn-auth/internal/repo/pgdb"
	"sn-auth/pkg/postgres"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (entity.User, error)
}

type Repositories struct {
	User
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: pgdb.NewUserRepo(pg),
	}
}
