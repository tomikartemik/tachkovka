package service

import (
	"tachkovka/internal/model"
	"tachkovka/internal/repository"
)

type RecordService struct {
	repo repository.Record
}

func NewRecordService(repo repository.Record) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) CreateRecord(record model.Record) error {
	return s.repo.CreateRecord(record)
}

func (s *RecordService) GetAllRecords() ([]model.Record, error) {
	return s.repo.GetAllRecords()
}

func (s *RecordService) GetRecordsByTabelNumber(tabelNumber int) ([]model.Record, error) {
	records, err := s.repo.GetRecordsByTabelNumber(tabelNumber)
	if err != nil {
		return nil, err
	}
	return records, nil
}
