package repositories

import (
	"Billboard/internal/models"
	"Billboard/pkg/database"
)

func GetScheduleByUID(UserID int) ([]models.Schedule, error) {
	var schedules []models.Schedule
	err := database.DB.Where("user_id = ?", UserID).Find(&schedules).Error
	return schedules, err
}

func UpdateSchedule(schedule models.Schedule) error {
	err := database.DB.Where("id = ?", schedule.ID).Updates(&schedule).Error
	return err
}
