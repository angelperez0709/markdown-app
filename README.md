# Markdown Notes API

Simple API for uploading Markdown, rendering HTML, listing notes, and checking grammar.

Based on this roadmap project: https://roadmap.sh/projects/markdown-note-taking-app

## Requirements
- Go 1.25+
- PostgreSQL

## Configuration
Set these environment variables (see [`config.Load`](internal/config/config.go) in [internal/config/config.go]):
- `APP_ADDR` (e.g. `8080`)
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

## Run
```sh
go run ./cmd/markdown-app
```

## API Endpoints
Base path: `/api/v1` (see [`routers.SetupRoutes`](internal/routers/router.go) in [internal/routers/router.go](internal/routers/router.go)).

### `POST /api/v1/notes`
Upload a Markdown file.

- Form fields:
  - `title` (optional)
  - `markdown` (file)

Response: JSON with the new note ID.

Handler: [`handlers.MakeUploadMarkdownHandler`](internal/handlers/markDownHandlers.go) in [internal/handlers/markDownHandlers.go](internal/handlers/markDownHandlers.go)

### `GET /api/v1/notes`
List all notes (id + title).

Handler: [`handlers.ListAllMarkdownsHandler`](internal/handlers/markDownHandlers.go) in [internal/handlers/markDownHandlers.go](internal/handlers/markDownHandlers.go)

### `GET /api/v1/notes/{id}/render`
Return rendered HTML for a note.

Handler: [`handlers.GetHTMLHandler`](internal/handlers/markDownHandlers.go) in [internal/handlers/markDownHandlers.go](internal/handlers/markDownHandlers.go)

### `POST /api/v1/notes/check`
Check grammar on a Markdown file.

- Form fields:
  - `markdown` (file)

Handler: [`handlers.CheckGrammarHandler`](internal/handlers/markDownHandlers.go) in [internal/handlers/markDownHandlers.go](internal/handlers/markDownHandlers.go)

## Notes Storage
Notes are saved to PostgreSQL by [`repository.MarkdownRepository`](internal/repository/markDownRepository.go) in [internal/repository/markDownRepository.go](internal/repository/markDownRepository.go), and Markdown files are stored under `markdowns/` by [`services.MarkdownService`](internal/services/markDownService.go) in [internal/services/markDownService.go](internal/services/markDownService.go).