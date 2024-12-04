package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
	"errors"
	"gorm.io/gorm"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func GetUserByName(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(ID int) (*models.User, error) {
	var user models.User
	err := database.DB.Where("user_id = ?", ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUser(username string) (bool, error) {
	var user models.User
	err := database.DB.Where("user_name = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	}

	return false, err

}

func AddUser(user *models.User) error {
	err := database.DB.Create(&user).Error
	return err
}

func UpdateUser(user *models.User) error {
	err := database.DB.Save(&user).Error
	return err
}

func DeleteUserByName(username string) error {
	err := database.DB.Where("user_name = ?", username).Delete(&models.User{}).Error
	return err
}

func DeleteUserByID(ID int) error {
	err := database.DB.Where("user_id = ?", ID).Delete(&models.User{}).Error
	return err
}
