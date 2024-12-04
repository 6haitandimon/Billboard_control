package models

type User struct {
	ID               int    `gorm:"primaryKey;autoIncrement;column:user_id"`
	FullName         string `gorm:"size:100;not null"`
	Phone            string `gorm:"size:15;not null"`
	RegistrationDate string `gorm:"not null"`
	DeviceSerial     string `gorm:"size:50;not null"`
	PasswordHash     string `gorm:"size:255;not null"`
	RoleID           int
	UserName         string `gorm:"size:50;not null"`
}

//ALTER TABLE Users CHANGE COLUMN device_serial_number device_serial VARCHAR(50) NOT NULL;
