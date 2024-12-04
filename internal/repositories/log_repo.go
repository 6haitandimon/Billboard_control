package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetLogs(user *models.User) ([]models.Log, error) {
	var Logs []models.Log
	err := database.DB.Where("user_id = ?", user.ID).Find(&Logs).Error
	return Logs, err
}
