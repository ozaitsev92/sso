package models

import "time"

type App struct {
	ID        int
	Name      string
	Secret    string
	CreatedAt time.Time
}
