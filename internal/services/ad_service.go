package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

func GetScheduleByUserID(UserID int) ([]models.Schedule, error) {
	schedule, err := repositories.GetScheduleByUID(UserID)
	if err != nil {
		return schedule, err
	}
	return schedule, nil
}

func UpdateSchedule(schedule models.Schedule) error {
	err := repositories.UpdateSchedule(schedule)
	if err != nil {
		return err
	}
	return nil
}

func DeserializeAdIDs(adIDsStr string) ([]int, error) {
	var adIDs []int
	err := json.Unmarshal([]byte(adIDsStr), &adIDs)
	if err != nil {
		return nil, err
	}
	return adIDs, nil
}

func UpdateStatistic(DeviceID int, adId int) error {
	stat, err := repositories.GetStatByADSId(DeviceID, adId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var statistic = models.AdStatistics{
			DeviceID:     DeviceID,
			AdID:         adId,
			DisplayCount: 1,
			LastUpdated:  time.Now(),
		}
		err = repositories.AddStatistic(statistic)
	} else if err != nil {
		return err
	}

	stat.DisplayCount = stat.DisplayCount + 1
	stat.LastUpdated = time.Now()

	err = repositories.UpdateStatistic(stat)
	if err != nil {
		return err
	}
	return nil

}

func GetStatisticByADSId(adId int) ([]models.AdStatistics, error) {
	statistics, err := repositories.GetStatsByADSId(adId)
	if err != nil {
		return statistics, err
	}
	return statistics, nil
}
