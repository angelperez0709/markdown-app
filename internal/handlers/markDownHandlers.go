package handlers

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func CheckGrammarHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Get Markdown Handler"))
}
func ListAllMarkdownsHandler(writer http.ResponseWriter, request *http.Request) {

}
func UploadMarkdownHandler(writer http.ResponseWriter, request *http.Request) {

}

func GetNoteRenderHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	writer.Write([]byte("Get Note Render Handler for ID: " + id))

}