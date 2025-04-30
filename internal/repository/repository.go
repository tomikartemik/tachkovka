package repository

import (
	"gorm.io/gorm"
	"tachkovka/internal/model"
)

type Repository struct {
	Record
	Tables
	BreakdownClassification
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Record:                  NewRecordRepository(db),
		Tables:                  NewTablesRepository(db),
		BreakdownClassification: NewBreakdownClassification(db),
	}
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecordsByTabelNumber(tabelNumber int) ([]model.Record, error)
}

type Tables interface {
	GetVersions() ([]model.Version, error)
	GetTables(name string) ([]model.Table, error)
	GetVersionsTable() ([]model.Version, error)
}

type BreakdownClassification interface {
	GetClassifications() ([]model.BreakdownClassification, error)
}
