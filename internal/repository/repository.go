package repository

import (
	"gorm.io/gorm"
	"tachkovka/internal/model"
)

type Repository struct {
	Record
	Tables
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Record: NewRecordRepository(db),
		Tables: NewTablesRepository(db),
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
