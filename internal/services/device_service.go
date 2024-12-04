package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
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

/*func UpdateDevice(device models.Device) (models.Device, error) {

}
*/

func DeleteDevice(ID int) error {
	device, err := repositories.GetDeviceById(ID)
	if err != nil {
		return err
	}

	return repositories.DeleteDevice(&device)
}
