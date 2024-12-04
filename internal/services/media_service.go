package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
)

func GetAllMedia() ([]models.Advertisements, error) {
	return repositories.GetAllMedia()
}

func GetMediaByID(ID int) (models.Advertisements, error) {
	return repositories.GetMediaByID(ID)
}

func DeleteMedia(media *models.Advertisements) error {
	return repositories.DeleteMedia(&media)
}
