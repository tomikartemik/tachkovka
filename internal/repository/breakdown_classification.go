package repository

import (
	"gorm.io/gorm"
	"tachkovka/internal/model"
)

type BreakdownClassificationRepository struct {
	db *gorm.DB
}

func NewBreakdownClassification(db *gorm.DB) *BreakdownClassificationRepository {
	return &BreakdownClassificationRepository{db: db}
}

func (r *BreakdownClassificationRepository) GetClassifications() ([]model.BreakdownClassification, error) {
	var classifications []model.BreakdownClassification
	result := r.db.Find(&classifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return classifications, nil
}
