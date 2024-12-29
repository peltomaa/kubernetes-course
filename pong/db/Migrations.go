package db

import "database/sql"

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS pong_clicks (
    id SERIAL PRIMARY KEY
  );`

	_, err := db.Exec(query)

	return err
}
