// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package postgres

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
