package utils

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandlePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return fmt.Errorf("failed to handle PostgreSQL error: %w", pgErr)
	}
	return fmt.Errorf("failed to handle error: %w", err)
}
