package service

import (
	"sort"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

type CampaignService struct {
	campaignRepository repo.Campaign
	clickRepository    repo.Click
}

func NewCampaignService(repos *repo.Repository) *CampaignService {
	return &CampaignService{
		campaignRepository: repos.Campaign,
		clickRepository:    repos.Click,
	}
}

func (s *CampaignService) GetCampaignActivity(campaignID int64, countHours int64) (*model.ActivityMetrics, error) {
	campaign, err := s.campaignRepository.GetByID(campaignID)
	if err != nil {
		return nil, err
	}

	metrics := &model.ActivityMetrics{
		CampaignID:   campaignID,
		CreatedAt:    campaign.CreatedAt,
		HourlyClicks: make([]*model.HourlyClicks, countHours),
	}

	// Calculate time range for hours
	startTime := campaign.CreatedAt
	
	for i := 0; i < int(countHours); i++ {
		metrics.HourlyClicks[i] = &model.HourlyClicks{
			Hour: startTime.Add(time.Duration(i) * time.Hour).Format("15:04"),
		}
	}

	clicks := s.clickRepository.GetByCampaignID(campaignID)

	// Count clicks within specified hours
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

		if clickDateTime.Before(startTime) || clickDateTime.After(startTime.Add(time.Duration(countHours) * time.Hour)) {
			continue
		}

		metrics.TotalClicks++
		hourIndex := int(clickDateTime.Sub(startTime).Hours())
		if hourIndex >= 0 && hourIndex < int(countHours) {
			metrics.HourlyClicks[hourIndex].Clicks++
		}
	}

	return metrics, nil
}

func (s *CampaignService) GetAllCampaigns(page, perPage int) *model.CampaignList {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10 // значение по умолчанию
	}

	campaigns := s.campaignRepository.GetAllCampaigns()

	campaignsInfo := make([]model.CampaignInfo, 0, len(campaigns))
	
	// Собираем все кампании
	for _, campaign := range campaigns {
		campaignsInfo = append(campaignsInfo, model.CampaignInfo{
			ID:        campaign.ID,
			Name:      campaign.Name,
			CreatedAt: campaign.CreatedAt,
		})

	}

	// Сортируем по ID для консистентности
	sort.Slice(campaignsInfo, func(i, j int) bool {
		return campaignsInfo[i].ID < campaignsInfo[j].ID
	})

	// Вычисляем общее количество страниц
	total := len(campaigns)
	totalPages := (total + perPage - 1) / perPage

	// Вычисляем начальный и конечный индексы для текущей страницы
	start := (page - 1) * perPage
	end := start + perPage
	if end > total {
		end = total
	}
	if start > total {
		start = total
	}

	return &model.CampaignList{
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		Campaigns:  campaignsInfo[start:end],
	}
} 
