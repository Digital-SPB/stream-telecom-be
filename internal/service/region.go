package service

import (
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

type RegionService struct {
	regionRepository repo.Regions
	clickRepository repo.Click
}

func NewRegionService(repos *repo.Repository) *RegionService {
	return &RegionService{
		regionRepository: repos.Regions,
		clickRepository:  repos.Click,
	}
}

func (s *RegionService) GetMembersHeatMap(startDate, endDate time.Time) []*model.RegionMembersHeatMap {
	regions := s.regionRepository.GetAll()
	clicks := s.clickRepository.GetAll()

	// Карта регионов с множеством уникальных member_id
	// map[RegionID]map[MemberID]struct{}
	membersByRegion := make(map[int64]map[int64]struct{})

	// Инициализируем множества для каждого региона
	for _, region := range regions {
		membersByRegion[region.ID] = make(map[int64]struct{})
	}

	// Обрабатываем клики
	for _, click := range clicks {
		clickDateTime := time.Date(
			click.ClickDate.Year(),
			click.ClickDate.Month(),
			click.ClickDate.Day(),
			click.ClickTime.Hour(),
			click.ClickTime.Minute(),
			click.ClickTime.Second(),
			0,
			time.UTC,
		)

		// Пропускаем клики вне временного интервала
		if !startDate.IsZero() && clickDateTime.Before(startDate) {
			continue
		}
		if !endDate.IsZero() && clickDateTime.After(endDate) {
			continue
		}

		// Добавляем member_id в множество для соответствующего региона
		if _, exists := membersByRegion[click.RegionID]; exists {
			membersByRegion[click.RegionID][click.MemberID] = struct{}{}
		}
	}

	// Формируем результат
	res := make([]*model.RegionMembersHeatMap, 0)
	for _, region := range regions {
		res = append(res, &model.RegionMembersHeatMap{
			ID:           region.ID,
			Name:         region.Name,
			MembersCount: int64(len(membersByRegion[region.ID])), // количество уникальных пользователей
		})
	}

	return res
}
