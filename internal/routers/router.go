package routers

import (
	"github.com/angelperez0709/markdown-app/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/check-grammar", handlers.CheckGrammarHandler)
		r.Get("/notes/{id}/render", handlers.GetNoteRenderHandler)
		r.Get("/notes",handlers.ListAllMarkdownsHandler)
		
		r.Post("/notes", handlers.UploadMarkdownHandler)
	})

}
