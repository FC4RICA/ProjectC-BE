package data

import (
	"time"
)

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(name, email string) *Account {
	return &Account{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
}
