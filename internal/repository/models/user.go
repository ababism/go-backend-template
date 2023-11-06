package models

import (
	"github.com/google/uuid"
	"hw_3/internal/service/domain"
)

type User struct {
	Id       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Balance  float64   `db:"balance"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		Id:      u.Id,
		Name:    u.Username,
		Balance: u.Balance,
	}
}
