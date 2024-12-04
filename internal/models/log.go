package models

type Log struct {
	LogID           int    `gorm:"primaryKey;column:log_id" json:"log_id"`
	UserID          int    `gorm:"column:user_id" json:"user_id"`
	ActionType      string `gorm:"column:action_type" json:"action_type"`
	ActionTimestamp string `gorm:"column:action_timestamp;autoCreateTime" json:"action_timestamp"`
	DeviceID        int    `gorm:"column:device_id" json:"device_id"`
	Details         string `gorm:"column:details" json:"details"`
}
