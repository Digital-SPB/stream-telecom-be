package service

import (
	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Campaign interface {
	GetCampaignActivity(campaignID int64, countHours int64) (*model.ActivityMetrics, error)
	GetAllCampaigns(page, perPage int) *model.CampaignList
}

type Click interface {
}

type Regions interface {
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
