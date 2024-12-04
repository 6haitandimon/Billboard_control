package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
)

func GetLogsByUser(user *models.User) ([]models.Log, error) {
	return repositories.GetLogs(user)
}
