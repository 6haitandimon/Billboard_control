package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetAllDevices() ([]models.Device, error) {
	var devices []models.Device
	err := database.DB.Find(&devices).Error
	return devices, err
}

func GetDeviceById(id int) (models.Device, error) {
	var device models.Device
	err := database.DB.Where("device_id = ?", id).First(&device).Error
	return device, err
}

func AddDevice(device *models.Device) error {
	err := database.DB.Create(&device).Error
	return err
}

func DeleteDevice(device *models.Device) error {
	err := database.DB.Delete(&device).Error
	return err
}
