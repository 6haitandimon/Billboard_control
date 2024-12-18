package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"fmt"
)

func FetchAllDevices() ([]models.Device, error) {
	return repositories.GetAllDevices()
}

func CreateDevice(ID int) (models.Device, error) {
	var device models.Device
	device = models.Device{
		UserID:           ID,
		ConnectionStatus: false,
		LoadedAds:        "",
		GroupID:          4,
	}

	err := repositories.AddDevice(&device)

	return device, err
}

func UpdateDevice(device models.Device) (models.Device, error) {
	err := repositories.UpdateDeviceData(device)
	if err != nil {
		return device, err
	}
	updateDevice, err := repositories.GetDeviceById(device.DeviceID)
	if err != nil {
		return device, err
	}
	return updateDevice, nil
}

func GetUserDevices(ID int) ([]models.Device, error) {
	devices, err := repositories.GetDeviceByUserId(ID)
	if err != nil {
		return devices, err
	}
	return devices, nil
}

func GetDeviceByID(DeviceID int) (models.Device, error) {
	device, err := repositories.GetDeviceById(DeviceID)
	if err != nil {
		return device, err
	}
	return device, nil
}

func GetUserDevicesByGroup(UserID int, GroupId int) ([]models.Device, error) {
	devices, err := repositories.GetDeviceToGroup(UserID, GroupId)
	if err != nil {
		return devices, err
	}
	return devices, nil
}

func GetAllFreeDevices() ([]models.Device, error) {
	devices, err := repositories.GetDeviceByGroup(4)
	if err != nil {
		return devices, err
	}

	return devices, nil
}

func DeleteDevice(ID int) error {
	device, err := repositories.GetDeviceById(ID)
	if err != nil {
		return err
	}

	return repositories.DeleteDevice(&device)
}

func GetGroupDevice(Id int) ([]models.DeviceGroup, error) {
	var groups []models.DeviceGroup
	devices, err := repositories.GetDeviceByUserId(Id)
	if err != nil {
		return groups, err
	}

	groupIDsMap := make(map[int]struct{})
	for _, device := range devices {
		groupIDsMap[device.GroupID] = struct{}{}
	}

	for groupID := range groupIDsMap {
		//var group models.Group
		group, err := repositories.GetGroup(groupID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		groups = append(groups, group)
	}

	return groups, nil
}
