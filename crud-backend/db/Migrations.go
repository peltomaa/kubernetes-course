package db

import "database/sql"

func CreateTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    task TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
  );`

	if _, err := db.Exec(createTableQuery); err != nil {
		return err
	}

	addUpdatedAtColumnQuery := `
    ALTER TABLE tasks ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT NULL;
	`
	if _, err := db.Exec(addUpdatedAtColumnQuery); err != nil {
		return err
	}

	addIsDoneColumnQuery := `
    ALTER TABLE tasks ADD COLUMN IF NOT EXISTS is_done BOOLEAN DEFAULT FALSE;
	`

	if _, err := db.Exec(addIsDoneColumnQuery); err != nil {
		return err
	}

	return nil
}
