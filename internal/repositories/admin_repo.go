package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetAdminByUsername(username string) (*models.Administrator, error) {
	var admin models.Administrator
	err := database.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
