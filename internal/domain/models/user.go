package models

import "time"

type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	IsAdmin      bool
	CreatedAt    time.Time
}
