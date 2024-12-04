package models

type Advertisements struct {
	ID        int    `gorm:"primaryKey;column:ad_id;autoIncrement" json:"media_id"`
	MediaName string `gorm:"column:ad_name" json:"media_name"`
	FilePath  string `gorm:"column:ad_link" json:"media_path"`
}
