package config
import "time"

type Config struct {
	Addr string
	DB   DBConfig
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

type DBConfig struct {
}
