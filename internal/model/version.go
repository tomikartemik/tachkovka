package model

type Version struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}
