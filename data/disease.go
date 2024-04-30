package data

import "time"

type Disease struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Discription string    `json:"discription"`
	CreatedAt   time.Time `json:"created_at"`
}
