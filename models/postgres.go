package models

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open will open a SQL connection with the provided
// Postgres database. Callers of Open need to ensure
// that the connection is eventually closed via the
// db.Close() method.
func Open(cfg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "EmoryRoland",
		Password: "EmoryRoland",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (c PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}
