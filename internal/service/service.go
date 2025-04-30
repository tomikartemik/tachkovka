package service

import (
	"tachkovka/internal/model"
	"tachkovka/internal/repository"
)

type Service struct {
	Record
	Tables
	BreakdownClassification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Record:                  NewRecordService(repos.Record),
		Tables:                  NewTablesService(repos.Tables),
		BreakdownClassification: NewBreakdownClassificationService(repos.BreakdownClassification),
	}
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecordsByTabelNumber(tabelNumber int) ([]model.Record, error)
}

type Tables interface {
	GetTable(name string) ([]model.Table, error)
	GetVersionsTable() ([]model.Version, error)
}

type BreakdownClassification interface {
	GetClassifications() ([]model.BreakdownClassification, error)
}
