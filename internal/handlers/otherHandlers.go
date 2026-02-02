package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CheckGrammarHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Get Markdown Handler"))
}



func GetNoteRenderHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	writer.Write([]byte("Get Note Render Handler for ID: " + id))
}

func GetHTMLHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	writer.Write([]byte("Get HTML Handler for ID: " + id))
}
