package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type DBService struct {
	db *sql.DB
}

func NewDBService(connectionString string) (*DBService, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	s := DBService{db: db}
	return &s, s.Ping()
}

func (s DBService) Ping() error {
	var result int
	err := s.db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		return err
	}
	log.Debugf("DB Ping Result: %d", result)
	return nil
}
