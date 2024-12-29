package repositories

import "database/sql"

type PongButtonClick struct{}

func (pb *PongButtonClick) Create(db *sql.DB) error {
	query := `INSERT INTO pong_clicks DEFAULT VALUES;`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (pb *PongButtonClick) Count(db *sql.DB) (int, error) {
	query := `SELECT COUNT(*) FROM pong_clicks;`

	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
