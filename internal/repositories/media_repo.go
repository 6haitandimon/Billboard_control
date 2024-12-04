package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetAllMedia() ([]models.Advertisements, error) {
	var mediaList []models.Advertisements
	err := database.DBADS.Find(&mediaList).Error
	if err != nil {
		return nil, err
	}
	return mediaList, nil
}

func GetMediaByID(ID int) (models.Advertisements, error) {
	var advertisement models.Advertisements
	err := database.DBADS.Where("media_id = ?", ID).First(&advertisement).Error
	return advertisement, err
}

func DeleteMedia(media *models.Advertisements) error {
	err := database.DBADS.Delete(media).Error
	return err
}
