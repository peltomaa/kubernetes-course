package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	psqlURL := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable", user, pass, host, port)

	db, err := sql.Open("postgres", psqlURL)
	if err != nil {
		return nil, err
	}

	RunMigrations(db)

	return db, nil
}
