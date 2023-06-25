package auth

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TokenValue string    `json:"token_value"`
	Expiration time.Time `json:"expiration"`
}

type Session struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}
type Users struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
