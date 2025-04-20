package model

import "time"

type Record struct {
	ID          int       `gorm:"primaryKey;autoIncrement;uniqueIndex;not null" json:"id"`
	TabelNumber int       `gorm:"not null" json:"tabel_number"`
	Action      string    `gorm:"not null" json:"action"`
	MachineType string    `gorm:"not null" json:"machine_type"`
	Machine     string    `gorm:"not null" json:"machine"`
	Details     string    `gorm:"not null" json:"details"`
	DateTime    time.Time `gorm:"not null" json:"date_time"`
}
