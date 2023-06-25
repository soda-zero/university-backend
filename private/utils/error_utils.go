package utils

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandlePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Println(pgErr)
		return fmt.Errorf("SqlError: %w", pgErr)
	}
	return fmt.Errorf("Error: %w", err)
}
