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
