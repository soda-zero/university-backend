package auth

import (
	"context"
	"fmt"
	"time"
	"zeroCalSoda/university-backend/private/db"
	"zeroCalSoda/university-backend/private/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	pool *pgxpool.Pool
}

func (r *AuthRepository) Close() {
	r.pool.Close()
}
func NewAuthRepository() (*AuthRepository, error) {
	pool, err := db.ConnectPg()
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %w", err)
	}

	return &AuthRepository{
		pool: pool,
	}, nil
}
func (r *AuthRepository) Signup(userData Users) error {
	query := "INSERT INTO users(username, password) VALUES ($1, $2)"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 8)
	if err != nil {
		return err
	}

	if _, err := r.pool.Exec(context.Background(), query, userData.Username, string(hashedPassword)); err != nil {
		return utils.HandlePgError(err)
	}

	return nil
}

func (r *AuthRepository) Signin(username string) (Session, error) {
	var storedUser Users
	query := "SELECT id FROM users WHERE username = $1"
	err := r.pool.QueryRow(context.Background(), query, username).Scan(&storedUser.ID)
	if err != nil {
		return Session{}, utils.HandlePgError(err)
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessionQuery := "INSERT INTO session(user_id, token, expiration) VALUES ($1, $2, $3)"
	_, err = r.pool.Exec(context.Background(), sessionQuery, storedUser.ID, sessionToken, expiresAt)
	if err != nil {
		return Session{}, utils.HandlePgError(err)
	}

	sessionData := Session{
		Token:      sessionToken,
		Expiration: expiresAt,
	}

	return sessionData, nil
}
