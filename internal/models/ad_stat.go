package models

import "time"

type AdStatistics struct {
	StatID       int       `gorm:"primaryKey;column:stat_id" json:"stat_id"`                     // Уникальный идентификатор статистики
	DeviceID     int       `gorm:"not null;column:device_id" json:"device_id"`                   // Привязка к таблице устройств
	AdID         int       `gorm:"not null;column:ad_id" json:"ad_id"`                           // Привязка к таблице рекламы
	DisplayCount int       `gorm:"not null;default:0;column:display_count" json:"display_count"` // Количество показов
	LastUpdated  time.Time `gorm:"autoUpdateTime:milli;column:last_updated" json:"last_updated"` // Время последнего обновления
}

func (AdStatistics) TableName() string {
	return "AdStatistics"
}
