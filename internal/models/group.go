package models

type DeviceGroup struct {
	GroupID   int    `gorm:"primaryKey;column:group_id" json:"group_id"`
	GroupName string `gorm:"column:group_name" json:"group_name"`
	UserID    int    `gorm:"column:user_id" json:"user_id"`
}

func (DeviceGroup) TableName() string {
	return "DeviceGroups"
}
