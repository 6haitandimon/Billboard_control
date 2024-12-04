package services

import (
	"log"
	"time"

	"Billboard/internal/models"
	"Billboard/internal/repositories"
)

type AdService struct {
	ScheduleRepo repositories.ScheduleRepository
}

func NewAdService(scheduleRepo repositories.ScheduleRepository) *AdService {
	return &AdService{ScheduleRepo: scheduleRepo}
}

func (s *AdService) StartAdDispatch(userID int) {
	go func() {
		for {
			// Получить расписание для пользователя
			schedules, err := s.ScheduleRepo.GetSchedulesForUser(userID)
			if err != nil {
				log.Printf("Error fetching schedules for user %d: %v\n", userID, err)
				time.Sleep(time.Minute) // Задержка перед повторной попыткой
				continue
			}

			// Обработка расписания
			for _, schedule := range schedules {
				nextShowTime := calculateNextShowTime(schedule.DisplayFrequencyPerHour)
				log.Printf("Next show for user %d at %s\n", userID, nextShowTime)

				// Ждём следующего показа
				time.Sleep(time.Until(nextShowTime))

				// Логика отправки рекламы
				ad, err := s.ScheduleRepo.GetAd(schedule.AdID)
				if err != nil {
					log.Printf("Error fetching ad %d: %v\n", schedule.AdID, err)
					continue
				}

				s.sendAdToFront(userID, ad)
			}
		}
	}()
}

func calculateNextShowTime(frequency int) time.Time {
	interval := time.Hour / time.Duration(frequency)
	return time.Now().Add(interval)
}

func (s *AdService) sendAdToFront(userID int, ad models.Ad) {
	// Здесь нужно реализовать логику отправки рекламы на фронт.
	log.Printf("Sending ad %d to user %d\n", ad.ID, userID)
}
