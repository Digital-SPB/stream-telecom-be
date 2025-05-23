package service

import (
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Campaign interface {
	GetCampaignActivity(campaignID int64, countHours int64) (*model.ActivityMetrics, error)
	GetAllCampaigns(page, perPage int) *model.CampaignList
	GetCreationDynamic(start time.Time, end time.Time, intervalType string) ([]*model.IntervalResult, error)
}

type Click interface {
	GetClickDynamic(id int64) (*model.CampaignStats, error)
	GetCustomerReactionTime(campaignID int64) (*model.ReactionTimeMetrics, error)
	GetTimeActivity() *model.TimeActivityResponse
	GetDailyTimeActivity(targetDate time.Time) *model.DailyTimeActivityResponse
}

type Regions interface {
	GetMembersHeatMap(startDate, endDate time.Time) []*model.RegionMembersHeatMap
	GetRegionsInfo() []*model.RegionInfo
	GetCountClick(startDate, endDate time.Time) []*model.CountClickByRegion
}

type Service struct {
	Campaign
	Click
	Regions
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Click:    NewClickService(repos),
		Campaign: NewCampaignService(repos),
		Regions:  NewRegionService(repos),
	}
}
