package model

type BreakdownClassification struct {
	ID     int    `gorm:"auto_increment;primary_key" json:"id"`
	Type   string `json:"type"`
	Reason string `json:"reason"`
}
