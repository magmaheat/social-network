package repo

import (
	"context"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/entity"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/repo/pgdb"
	"github.com/magmaheat/social-network/tree/main/sn-auth/pkg/postgres"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUserByUsernameAndPassword(ctx context.Context, username string) (entity.User, error)
	GetUserById(ctx context.Context, id int) (entity.User, error)
}

type Repositories struct {
	User
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: pgdb.NewUserRepo(pg),
	}
}
