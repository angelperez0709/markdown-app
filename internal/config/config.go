package config

import (
	"os"
	"time"
)

type Config struct {
	Addr         string
	DB           DBConfig
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() Config {
	db := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	
	c := Config{
		Addr:         ":" + os.Getenv("APP_ADDR"),
		DB:           db,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
		IdleTimeout:  time.Minute * 2,
	}
	return c
}
