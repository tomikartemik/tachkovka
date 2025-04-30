package model

type Eto struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}
