package service

import (
	"tachkovka/internal/model"
	"tachkovka/internal/repository"
)

type Service struct {
	Record
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Record: NewRecordService(repos.Record),
	}
}

type Record interface {
	CreateRecord(record model.Record) error
	GetAllRecords() ([]model.Record, error)
	GetRecordsByTabelNumber(tabelNumber int) ([]model.Record, error)
}
