package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string, logger *slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("Database connected successfully")
	return db, nil
}
