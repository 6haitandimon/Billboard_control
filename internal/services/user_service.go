package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
)

func FetchAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUser(ID int) (*models.User, error) {
	return repositories.GetUserByID(ID)
}

func DeleteUser(ID int) error {
	return repositories.DeleteUserByID(ID)
}
