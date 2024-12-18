package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetStatByADSId(DeviceID int, AdsID int) (models.AdStatistics, error) {
	var ads models.AdStatistics
	err := database.DB.Where("ad_id = ? AND device_id = ?", AdsID, DeviceID).First(&ads).Error
	if err != nil {
		return ads, err
	}
	return ads, nil
}

func GetStatsByADSId(AdsID int) ([]models.AdStatistics, error) {
	var ads []models.AdStatistics
	err := database.DB.Where("ad_id = ?", AdsID).Find(&ads).Error
	if err != nil {
		return ads, err
	}
	return ads, nil
}

func AddStatistic(Statistic models.AdStatistics) error {
	err := database.DB.Create(&Statistic).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatistic(Statistic models.AdStatistics) error {
	err := database.DB.Where("stat_id = ?", Statistic.StatID).Updates(&Statistic).Error
	if err != nil {
		return err
	}
	return nil
}
