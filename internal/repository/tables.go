package repository

import (
	"gorm.io/gorm"
	"tachkovka/internal/model"
)

type TablesRepository struct {
	db *gorm.DB
}

func NewTablesRepository(db *gorm.DB) *TablesRepository {
	return &TablesRepository{db: db}
}

func (r *TablesRepository) GetTableVersion() ([]model.Version, error) {
	versions := []model.Version{}
	err := r.db.Find(&versions).Error
	return versions, err
}

func (r *TablesRepository) GetVersions() ([]model.Version, error) {
	versions := []model.Version{}
	err := r.db.Find(&versions).Error
	return versions, err
}

func (r *TablesRepository) GetTables(name string) ([]model.Table, error) {
	tables := []model.Table{}
	switch name {
	case "excavator":
		err := r.db.Model(&model.Excavator{}).Find(&tables).Error
		return tables, err
	case "dump_truck":
		err := r.db.Model(&model.DumpTruck{}).Find(&tables).Error
		return tables, err
	case "down_time":
		err := r.db.Model(&model.DownTime{}).Find(&tables).Error
		return tables, err
	case "type_of_work":
		err := r.db.Model(&model.TypeOfWork{}).Find(&tables).Error
		return tables, err
	case "work_place":
		err := r.db.Model(&model.WorkPlace{}).Find(&tables).Error
		return tables, err
	}
	return tables, nil
}
