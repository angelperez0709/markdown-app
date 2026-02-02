package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type MarkdownRepository struct {
	db *sql.DB
}

func NewMarkdownRepository(db *sql.DB) *MarkdownRepository {
	return &MarkdownRepository{db: db}
}

func (r *MarkdownRepository) SaveNote(title string, filePath string, htmlContent []byte, createdAt time.Time) (int64, error) {
	query := `INSERT INTO notes (title, file_path, html_content, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	err := r.db.QueryRow(query, title, filePath, htmlContent, createdAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert note: %w", err)
	}

	return id, nil
}

func (r *MarkdownRepository) GetAllNotes() ([]Note, error) {
	query := `SELECT id, title FROM notes`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.Id, &note.Title); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return notes, nil
}

func (r *MarkdownRepository) GetHTMLContentByID(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id cannot be empty")
	}

	query := `SELECT html_content FROM notes WHERE id = $1`
	var htmlContent string
	err := r.db.QueryRow(query, id).Scan(&htmlContent)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("note with id %s not found", id)
	}
	if err != nil {
		return "", fmt.Errorf("failed to get HTML content: %w", err)
	}
	return htmlContent, nil
}

type Note struct {
	Id           int64
	Title        string
	file_path    string
	html_content []byte
	created_at   time.Time
}
