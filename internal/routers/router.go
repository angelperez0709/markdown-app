package routers

import (
	"github.com/angelperez0709/markdown-app/internal/handlers"
	"github.com/angelperez0709/markdown-app/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, service *services.MarkdownService) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/notes/{id}/render", handlers.GetHTMLHandler(service))
		r.Get("/notes",handlers.ListAllMarkdownsHandler(service))
		r.Post("/notes", handlers.MakeUploadMarkdownHandler(service))
		r.Post("/notes/check", handlers.CheckGrammarHandler(service))
	})

}
