package repository

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
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

	modelMap := map[string]interface{}{
		"version":      model.Version{},
		"excavator":    model.Excavator{},
		"dump_truck":   model.DumpTruck{},
		"down_time":    model.DownTime{},
		"type_of_work": model.TypeOfWork{},
		"work_place":   model.WorkPlace{},
	}

	modelInstance, ok := modelMap[strings.ToLower(name)]
	if !ok {
		return tables, nil
	}

	// Создаем слайс нужного типа через рефлексию
	sliceType := reflect.SliceOf(reflect.TypeOf(modelInstance))
	slice := reflect.New(sliceType).Interface()

	err := r.db.Model(modelInstance).Find(slice).Error
	if err != nil {
		return nil, err
	}

	// Конвертируем результаты
	result := reflect.ValueOf(slice).Elem()
	for i := 0; i < result.Len(); i++ {
		item := result.Index(i)
		// Если item является указателем, получаем значение
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		id := item.FieldByName("ID").Int()
		name := item.FieldByName("Name").String()
		tables = append(tables, model.Table{
			ID:   int(id),
			Name: name,
		})
	}

	return tables, nil
}

func (r *TablesRepository) GetVersionsTable() ([]model.Version, error) {
	versions := []model.Version{}
	err := r.db.Find(&versions).Error
	return versions, err
}
