package storage

import (
	"database/sql"
	"time"

	"github.com/Webblurt/garantex-usdt-rates/internal/rates"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// ping to check connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveRate(r *rates.Rate) error {
	_, err := s.db.Exec(
		"INSERT INTO rates (ask, bid, timestamp) VALUES ($1, $2, $3)",
		r.Ask, r.Bid, r.Timestamp.Format(time.RFC3339),
	)
	return err
}
