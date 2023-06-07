package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"zeroCalSoda/university-backend/private/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))

	var greeting string
	db, _ := db.ConnectPg()
	err := db.QueryRow(context.Background(), "SELECT 'Hello, world!!'").Scan(&greeting)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}
	fmt.Println(greeting)
}
