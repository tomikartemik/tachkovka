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

	// Создаем карту соответствий имен таблиц моделям
	modelMap := map[string]interface{}{
		"version":      &model.Version{},
		"excavator":    &model.Excavator{},
		"dump_truck":   &model.DumpTruck{},
		"down_time":    &model.DownTime{},
		"type_of_work": &model.TypeOfWork{},
		"work_place":   &model.WorkPlace{},
	}

	// Приводим имя к нижнему регистру для унификации
	name = strings.ToLower(name)

	// Получаем модель из карты
	modelInstance, ok := modelMap[name]
	if !ok {
		return tables, nil
	}

	// Используем рефлексию для создания слайса нужного типа
	sliceType := reflect.SliceOf(reflect.TypeOf(modelInstance))
	slice := reflect.New(sliceType).Interface()

	// Выполняем запрос
	err := r.db.Model(modelInstance).Find(slice).Error
	if err != nil {
		return nil, err
	}

	// Конвертируем результат в []model.Table
	result := reflect.ValueOf(slice).Elem()
	for i := 0; i < result.Len(); i++ {
		item := result.Index(i).Interface()
		id := reflect.ValueOf(item).FieldByName("ID").Int()
		name := reflect.ValueOf(item).FieldByName("Name").String()
		tables = append(tables, model.Table{
			ID:   int(id),
			Name: name,
		})
	}

	return tables, nil
}
