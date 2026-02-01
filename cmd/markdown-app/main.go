package main

import (
	"database/sql"
	"log"

	"github.com/angelperez0709/markdown-app/internal/app"
	"github.com/angelperez0709/markdown-app/internal/config"
	"github.com/angelperez0709/markdown-app/internal/repository"
	"github.com/angelperez0709/markdown-app/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db := initializeDb(cfg.DB)
	repository := initializeRepository(db)
	service := initializeService(repository)

	app := app.Application{
		Config:  cfg,
		Service: service,
	}
	handler := app.Mount()
	if err := app.Run(handler); err != nil {
		log.Printf("An error ocurred %s", err.Error())
		panic(err)
	}

}

func initializeDb(dbConfig config.DBConfig) *sql.DB {
	db, err := repository.InitializeDbConnection(dbConfig)
	if err != nil {
		log.Printf("Failed to connect to database: %s", err.Error())
		panic(err)
	}
	return db
}
func initializeRepository(db *sql.DB) *repository.MarkdownRepository {
	repo := repository.NewMarkdownRepository(db)
	return repo
}

func initializeService(repo *repository.MarkdownRepository) *services.MarkdownService {
	service := services.NewMarkdownService(repo)
	return service
}
