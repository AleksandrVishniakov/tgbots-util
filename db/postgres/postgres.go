package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Configs struct {
	Host     string
	Port     int
	User string
	DBName   string
	Password string
	SSLMode  string
}

func ConnectionString(cfg Configs) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func MustPostgresDB(cfg Configs) *sql.DB {
	db, err := sql.Open("postgres", ConnectionString(cfg))
	if err != nil {
		log.Fatalf("Failed to open database connection: %s\n", err.Error())
	}

	return db
}