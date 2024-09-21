package pgdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/entity"
	"github.com/magmaheat/social-network/tree/main/sn-auth/internal/repo/repoerrs"
	"github.com/magmaheat/social-network/tree/main/sn-auth/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) CreateUser(ctx context.Context, user entity.User) (int, error) {
	sql, args, _ := u.Builder.
		Insert("users").
		Columns("username", "password").
		Values(user.Username, user.Password).
		Suffix("RETURNING id").
		ToSql()

	var id int
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
			return 0, fmt.Errorf("UserRepo.CreateUser - u.Pool.QueryRow: %v", err)
		}
	}

	return id, nil
}

func (u *UserRepo) GetUserByUsernameAndPassword(ctx context.Context, username string) (entity.User, error) {
	sql, args, _ := u.Builder.
		Select("id, username, password, created_at").
		From("users").
		Where("username = ?", username).
		ToSql()

	var user entity.User
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repoerrs.ErrNotFound
		}
		return entity.User{}, fmt.Errorf("UserRepo.GetUserByIdUsernameAndPassword = u.Pool.QueryRow: %v", err)
	}

	return user, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, id int) (entity.User, error) {
	sql, args, _ := u.Builder.
		Select("id, username, password, created_at").
		From("users").
		Where("id = ?", id).
		ToSql()

	var user entity.User
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repoerrs.ErrNotFound
		}

		return entity.User{}, fmt.Errorf("UserRepo.GetUserById = u.Pool.QueryRow: %v", err)
	}

	return user, nil
}
