package helpers

import (
	"database/sql"
	"fmt"
	"os"
)

func PostgreHelper() *sql.DB {
	connStr := os.Getenv("POSTGRE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("Database Error: %v\n", err)
	}

	return db
}