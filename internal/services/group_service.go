package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"fmt"
)

func CreateGroup(GroupName string, UserID int) ([]models.DeviceGroup, error) {
	var group = models.DeviceGroup{
		GroupName: GroupName,
		UserID:    UserID,
	}
	var groups []models.DeviceGroup
	err := repositories.CreateGroup(group)
	if err != nil {
		fmt.Println(err)
		return groups, err
	}

	//group, err = repositories.GetGroupByName(GroupName)

	groups, err = repositories.GetGroupByUserID(UserID)
	if err != nil {
		fmt.Println(err)
		return groups, err
	}
	return groups, nil
}

func AddToGroup(UserID int, DeviceID int, GroupID int) (models.Device, error) {
	device, err := repositories.GetDeviceById(DeviceID)
	if err != nil {
		return models.Device{}, err
	}

	device.UserID = UserID
	device.GroupID = GroupID

	err = repositories.UpdateDeviceData(device)

	if err != nil {
		return models.Device{}, err
	}
	return device, nil

}

func DeleteOnGroup(DeviceID int, GroupID int) error {
	device, err := repositories.GetDeviceById(DeviceID)
	if err != nil {
		return err
	}

	device.GroupID = 1
	device.DeviceID = DeviceID

	err = repositories.UpdateDeviceData(device)
	if err != nil {
		return err
	}
	return nil
}
