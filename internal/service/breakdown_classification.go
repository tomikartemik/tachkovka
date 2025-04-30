package service

import (
	"tachkovka/internal/model"
	"tachkovka/internal/repository"
)

type BreakdownClassificationService struct {
	repo repository.BreakdownClassification
}

func NewBreakdownClassificationService(repo repository.BreakdownClassification) *BreakdownClassificationService {
	return &BreakdownClassificationService{repo: repo}
}

func (s *BreakdownClassificationService) GetClassifications() ([]model.BreakdownClassification, error) {
	return s.repo.GetClassifications()
}
