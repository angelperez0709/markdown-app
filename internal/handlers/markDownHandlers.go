package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/angelperez0709/markdown-app/internal/services"
)

func MakeUploadMarkdownHandler(service *services.MarkdownService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		title := request.FormValue("title")

		file, _, err := request.FormFile("markdown")
		if err != nil {
			http.Error(writer, "Markdown file is required", http.StatusBadRequest)
			return
		}
		defer file.Close()

		markdownContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(writer, "Error reading file", http.StatusInternalServerError)
			return
		}

		noteID, err := service.SaveNote(title, markdownContent)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)

		writer.Write([]byte(`{"id": ` + string(rune(noteID)) + `, "message": "Note saved successfully"}`))
	}
}

func ListAllMarkdownsHandler(service *services.MarkdownService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := service.ListAllNotes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert notes to JSON bytes
		jsonData, err := json.Marshal(notes)
		if err != nil {
			http.Error(w, "Failed to serialize notes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
