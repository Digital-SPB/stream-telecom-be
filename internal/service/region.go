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

func (s *RegionService) GetHeatMap(startDate, endDate time.Time) []*model.RegionHeatMap {
	regions := s.regionRepository.GetAll()
	clicks := s.clickRepository.GetAll()

	heatMap := make(map[int64]int64)

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

		if !startDate.IsZero() && clickDateTime.Before(startDate) {
			continue
		}
		if !endDate.IsZero() && clickDateTime.After(endDate) {
			continue
		}

		heatMap[click.RegionID]++
	}

	res := make([]*model.RegionHeatMap, 0)

	for _, region := range regions {
		res = append(res, &model.RegionHeatMap{
			ID:         region.ID,
			Name:       region.Name,
			ClickCount: heatMap[region.ID],
		})
	}

	return res
}
