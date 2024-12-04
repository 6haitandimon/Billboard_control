package models

type Administrator struct {
	ID           int    `gorm:"primaryKey;column:admin_id"`
	FullName     string `gorm:"size:100;not null"`
	Username     string `gorm:"size:50;unique;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	RoleID       int
}
