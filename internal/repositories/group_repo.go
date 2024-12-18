package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetGroup(ID int) (models.DeviceGroup, error) {
	var group models.DeviceGroup
	err := database.DB.Where("group_id = ?", ID).First(&group).Error
	return group, err
}

func GetGroupByName(Name string) (models.DeviceGroup, error) {
	var group models.DeviceGroup
	err := database.DB.Where("group_name = ?", Name).First(&group).Error
	return group, err
}

func GetGroupByUserID(ID int) ([]models.DeviceGroup, error) {
	var group []models.DeviceGroup
	err := database.DB.Where("user_id = ?", ID).Find(&group).Error
	return group, err
}

func CreateGroup(group models.DeviceGroup) error {
	err := database.DB.Create(&group).Error
	return err
}

//func UpdateGroup(group models.DeviceGroup) error {
//
//}
