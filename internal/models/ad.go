package models

type Ad struct {
	ID       int    `gorm:"primaryKey;column:ad_id" json:"ad_id"`
	FilePath string `gorm:"column:file_path" json:"file_path"`
	Type     string `gorm:"column:type" json:"type"`
}
