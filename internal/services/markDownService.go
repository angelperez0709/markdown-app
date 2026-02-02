package services

import (
	"bytes"
	"fmt"
	"github.com/angelperez0709/markdown-app/internal/repository"
	"github.com/jdkato/prose/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"os"
	"path/filepath"
	"time"
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

func (s *MarkdownService) GetHTMLContentByID(id string) ([]byte, error) {
	htmlContent, err := s.repository.GetHTMLContentByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve HTML content: %w", err)
	}
	return []byte(htmlContent), nil
}

func (s *MarkdownService) CheckGrammar(markdownContent []byte) ([]GrammarIssue, error) {
	text := string(markdownContent)
	fmt.Println("Checking grammar for content:", text)

	// Create a new prose document
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %w", err)
	}

	var issues []GrammarIssue

	// Check for spelling errors
	for _, tok := range doc.Tokens() {
		// Skip punctuation and short words
		if len(tok.Text) < 3 || !isWord(tok.Text) {
			continue
		}

		// Check if word is misspelled (basic check)
		if !isValidWord(tok.Text) {
			issues = append(issues, GrammarIssue{
				Type:        "spelling",
				Text:        tok.Text,
				Suggestion:  getSuggestions(tok.Text),
				Description: fmt.Sprintf("Possible spelling error: '%s'", tok.Text),
			})
		}
	}

	for _, sent := range doc.Sentences() {
		words := countWords(sent.Text)

		if words > 40 {
			issues = append(issues, GrammarIssue{
				Type:        "readability",
				Text:        sent.Text,
				Suggestion:  []string{"Consider breaking this into shorter sentences"},
				Description: fmt.Sprintf("Sentence is too long (%d words). Consider splitting it.", words),
			})
		}

		// Check for passive voice
		if containsPassiveVoice(sent.Text) {
			issues = append(issues, GrammarIssue{
				Type:        "style",
				Text:        sent.Text,
				Suggestion:  []string{"Consider using active voice"},
				Description: "Sentence may be in passive voice",
			})
		}
	}

	// Check for repeated words
	wordCount := make(map[string]int)
	for _, tok := range doc.Tokens() {
		if isWord(tok.Text) {
			wordCount[tok.Text]++
		}
	}

	for word, count := range wordCount {
		if count > 5 && len(word) > 4 {
			issues = append(issues, GrammarIssue{
				Type:        "repetition",
				Text:        word,
				Suggestion:  []string{"Consider using synonyms"},
				Description: fmt.Sprintf("Word '%s' appears %d times. Consider varying vocabulary.", word, count),
			})
		}
	}

	return issues, nil
}

func countWords(text string) int {
	count := 0
	inWord := false
	for _, r := range text {
		if r == ' ' || r == '\t' || r == '\n' {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}
	return count
}

// Helper functions
func isWord(text string) bool {
	for _, r := range text {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

func isValidWord(word string) bool {
	// Simple validation - in production use a proper dictionary
	// For now, just check if it's all letters
	return isWord(word)
}

func getSuggestions(word string) []string {
	// In production, use a spell-checking library like go-aspell
	return []string{"No suggestions available"}
}

func containsPassiveVoice(sentence string) bool {
	// Simple passive voice detection
	// Look for "was", "were", "been", "being" + past participle
	passiveIndicators := []string{" was ", " were ", " been ", " being "}
	lowerSentence := " " + sentence + " "
	for _, indicator := range passiveIndicators {
		if contains(lowerSentence, indicator) {
			return true
		}
	}
	return false
}

func contains(text, substr string) bool {
	return len(text) >= len(substr) && (text == substr || len(text) > len(substr) &&
		(text[:len(substr)] == substr || contains(text[1:], substr)))
}

// GrammarIssue represents a grammar or style issue found in the text
type GrammarIssue struct {
	Type        string   `json:"type"`        // spelling, grammar, style, readability, repetition
	Text        string   `json:"text"`        // The problematic text
	Suggestion  []string `json:"suggestion"`  // Suggested corrections
	Description string   `json:"description"` // Description of the issue
}
