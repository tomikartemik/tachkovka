package service

import (
	"tachkovka/internal/model"
	"tachkovka/internal/repository"
)

type TablesService struct {
	repo repository.Tables
}

func NewTablesService(repo repository.Tables) *TablesService {
	return &TablesService{repo: repo}
}

func (s *TablesService) GetVersions() ([]model.Version, error) {
	return s.repo.GetVersions()
}

func (s *TablesService) GetTable(name string) (*model.Table, error) {
	return s.repo.GetTables(name)
}
