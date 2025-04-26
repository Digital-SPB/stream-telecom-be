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
}

type Click interface {
	GetCustomerReactionTime(campaignID int64) (*model.ReactionTimeMetrics, error)
}

type Regions interface {
	GetMembersHeatMap(startDate, endDate time.Time) []*model.RegionMembersHeatMap
	GetCountClick(startDate, endDate time.Time) []*model.CountClickByRegion
}

type Service struct {
	Campaign
	Click
	Regions
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Campaign: NewCampaignService(repos),
		Click:    NewClickService(repos),
		Regions:  NewRegionService(repos),
	}
}
