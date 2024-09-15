package pgdb

import (
	"context"
	"sn-auth/internal/entity"
	"sn-auth/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) CreateUser(ctx context.Context, user entity.User) (int, error) {
	return 0, nil
}

func (u *UserRepo) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}
