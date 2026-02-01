package services

import (
	"github.com/angelperez0709/markdown-app/internal/repository"
)

type GrammarService struct {
	repository *repository.MarkdownRepository
}


func NewGrammarService(repository *repository.MarkdownRepository) *GrammarService {
	return &GrammarService{
		repository: repository,
	}
}
