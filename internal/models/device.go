package models

type Device struct {
	DeviceID         int    `gorm:"primaryKey;column:device_id;autoIncrement" json:"device_id"`
	UserID           int    `gorm:"column:user_id" json:"user_id"`
	ConnectionStatus bool   `gorm:"column:connection_status;default:true" json:"connection_status"`
	LoadedAds        string `gorm:"column:loaded_ads" json:"loaded_ads"`
	GroupID          int    `gorm:"column:group_id" json:"group_id"`
}
