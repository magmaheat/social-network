package repo

import (
	"sn-auth/internal/repo/pgdb"
	"sn-auth/pkg/postgres"
)

type User interface{}

type Repositories struct {
	User
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: pgdb.NewUserRepo(pg),
	}
}
