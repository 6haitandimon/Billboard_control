package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"encoding/json"
)

func GetScheduleByUserID(UserID int) ([]models.Schedule, error) {
	schedule, err := repositories.GetScheduleByUID(UserID)
	if err != nil {
		return schedule, err
	}
	return schedule, nil
}

func DeserializeAdIDs(adIDsStr string) ([]int, error) {
	var adIDs []int
	err := json.Unmarshal([]byte(adIDsStr), &adIDs)
	if err != nil {
		return nil, err
	}
	return adIDs, nil
}
