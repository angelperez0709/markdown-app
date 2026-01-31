package app

import (
	"log"
	"net/http"
	"github.com/angelperez0709/markdown-app/internal/config"
	"github.com/angelperez0709/markdown-app/internal/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config config.Config
}

func (application Application) Run(h http.Handler) error {
	srv := &http.Server{
		Addr:         application.Config.Addr,
		Handler:      h,
		WriteTimeout: application.Config.WriteTimeout,
		ReadTimeout:  application.Config.ReadTimeout,
		IdleTimeout:  application.Config.IdleTimeout,
	}
	log.Printf("Server started at address %s", application.Config.Addr)

	return srv.ListenAndServe()
}
func (application Application) Mount() http.Handler {
	r := chi.NewRouter()
	setupMiddlewares(r)
	routers.SetupRoutes(r)
	return r
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
}
