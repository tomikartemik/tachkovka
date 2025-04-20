package repository

import (
	"gorm.io/gorm"
	"tachkovka/internal/model"
)

type Repository struct {
	Record
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Record: NewRecordRepository(db),
	}
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecordsByTabelNumber(tabelNumber int) ([]model.Record, error)
}
