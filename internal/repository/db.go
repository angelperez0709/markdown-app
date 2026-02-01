package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/angelperez0709/markdown-app/internal/config"
	_ "github.com/lib/pq"
)

func InitializeDbConnection(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return db, nil
}

func CloseDB(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
