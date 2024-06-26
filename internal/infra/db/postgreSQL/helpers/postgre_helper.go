package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func PostgreHelper() *sql.DB {
	connStr := os.Getenv("POSTGRE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("Database Error: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Print("failed to ping database: %w", err)
	}

	return db
}