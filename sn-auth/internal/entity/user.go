package entity

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"password"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
