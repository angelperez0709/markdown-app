package services

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/angelperez0709/markdown-app/internal/repository"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type MarkdownService struct {
	repository *repository.MarkdownRepository
}

func NewMarkdownService(repository *repository.MarkdownRepository) *MarkdownService {
	return &MarkdownService{
		repository: repository,
	}
}

func (s *MarkdownService) SaveNote(title string, markdownContent []byte) (int64, error) {
	if title == "" {
		title = "Untitled"
	}
	//save note in system
	filepath, err := s.saveMarkdownToFile(title, markdownContent)
	if err != nil {
		return 0, fmt.Errorf("failed to save markdown file: %w", err)
	}

	htmlContent, err := s.convertMarkdownToHTML(markdownContent)
	if err != nil {
		return 0, fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	// Save to repository
	noteID, err := s.repository.SaveNote(title, filepath, htmlContent, time.Now())
	if err != nil {
		return 0, fmt.Errorf("failed to save note: %w", err)
	}

	return noteID, nil
}

func (s *MarkdownService) saveMarkdownToFile(title string, content []byte) (string, error) {
	markdownFolder := "markdowns"
	if err := os.MkdirAll(markdownFolder, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	filename := fmt.Sprintf("%s.md", title)
    filePath := filepath.Join(markdownFolder, filename)

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return filePath, nil
}


func (s *MarkdownService) convertMarkdownToHTML(markdown []byte) ([]byte, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,           
			extension.Table,         
			extension.Strikethrough, 
			extension.Linkify,      
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), 
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(), 
			html.WithXHTML(), 
		),
	)

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := md.Convert(markdown, &buf); err != nil {
		return nil, err
	}

	// Sanitize HTML to prevent XSS attacks
	policy := bluemonday.UGCPolicy()

	// Allow additional safe elements for markdown
	policy.AllowAttrs("class").Matching(bluemonday.SpaceSeparatedTokens).OnElements("code", "span", "div")
	policy.AllowAttrs("id").Matching(bluemonday.SpaceSeparatedTokens).OnElements("h1", "h2", "h3", "h4", "h5", "h6")

	// Sanitize the HTML output
	sanitizedHTML := policy.SanitizeBytes(buf.Bytes())

	return sanitizedHTML, nil
}

func (s *MarkdownService) ListAllNotes() ([]repository.Note, error) {
	notes, err := s.repository.GetAllNotes()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve notes: %w", err)
	}
	return notes, nil
}