package repositories

import (
	"Billboard/internal/models"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{DB: db}
}

func (r *ScheduleRepository) GetSchedulesForUser(userID int) ([]models.Schedule, error) {
	var schedules []models.Schedule
	err := r.DB.Where("device_id IN (SELECT device_id FROM devices WHERE user_id = ?)", userID).
		Find(&schedules).Error
	return schedules, err
}

func (r *ScheduleRepository) GetAd(adID int) (models.Ad, error) {
	var ad models.Ad
	err := r.DB.First(&ad, "id = ?", adID).Error
	return ad, err
}
