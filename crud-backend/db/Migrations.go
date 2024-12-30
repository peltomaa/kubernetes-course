package db

import "database/sql"

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    task TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
  );`

	_, err := db.Exec(query)

	return err
}
