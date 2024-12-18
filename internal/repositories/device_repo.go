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

func GetDeviceByUserId(id int) ([]models.Device, error) {
	var device []models.Device
	err := database.DB.Where("user_id = ?", id).Find(&device).Error
	return device, err
}

func GetDeviceByGroup(Group int) ([]models.Device, error) {
	var device []models.Device
	err := database.DB.Where("group_id = ?", Group).Find(&device).Error
	return device, err
}

func GetDeviceToGroup(User int, Group int) ([]models.Device, error) {
	var device []models.Device
	err := database.DB.Where("user_id = ? AND group_id = ?", User, Group).Find(&device).Error
	return device, err
}

func UpdateDeviceData(DeviceData models.Device) error {
	err := database.DB.Model(&models.Device{}).Where("device_id = ?", DeviceData.DeviceID).Updates(&DeviceData).Error
	return err
}

func AddDevice(device *models.Device) error {
	err := database.DB.Create(&device).Error
	return err
}

func DeleteDevice(device *models.Device) error {
	err := database.DB.Delete(&device).Error
	return err
}
