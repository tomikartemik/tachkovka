package model

type EtoExcavator struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}

type EtoDumpTruck struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
}
