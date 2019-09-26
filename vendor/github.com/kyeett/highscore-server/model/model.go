package model

import (
	"time"

	"github.com/google/uuid"
)

// Score model
type Score struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Score     float64   `db:"score" json:"score"`
	User
	Game
}

type Game struct {
	Name string `db:"game_name" json:"game_name"`
}

type User struct {
	ID   string `db:"user_id" json:"user_id"`
	Name string `db:"user_name" json:"user_name"`
}
