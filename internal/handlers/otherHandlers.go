package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func GetNoteRenderHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	writer.Write([]byte("Get Note Render Handler for ID: " + id))
}