package models

type Schedule struct {
	ID      int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // Уникальный идентификатор
	GroupID int    `gorm:"not null;column:group_id" json:"group_id"`     // Связь с DeviceGroups
	UserID  int    `gorm:"not null;column:user_id" json:"user_id"`       // Связь с Users
	Freq    int64  `gorm:"not null;column:freq" json:"freq"`
	AdIDs   string `gorm:"column:ad_ids" json:"ad_ids"`
}

type ScheduleSender struct {
	ID      int   `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // Уникальный идентификатор
	GroupID int   `gorm:"not null;column:group_id" json:"group_id"`     // Связь с DeviceGroups
	UserID  int   `gorm:"not null;column:user_id" json:"user_id"`       // Связь с Users
	Freq    int64 `gorm:"not null;column:freq" json:"freq"`
	AdIDs   []int `gorm:"column:ad_ids" json:"ad_ids"`
}

func (Schedule) TableName() string {
	return "Schedule"
}
