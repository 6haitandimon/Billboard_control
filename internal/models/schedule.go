package models

type Schedule struct {
	ScheduleID              int `gorm:"primaryKey;column:schedule_id" json:"schedule_id"`
	AdID                    int `gorm:"column:ad_id" json:"ad_id"`
	DisplayFrequencyPerHour int `gorm:"column:display_frequency_per_hour" json:"display_frequency_per_hour"`
	DeviceID                int `gorm:"column:device_id" json:"device_id"`
}
