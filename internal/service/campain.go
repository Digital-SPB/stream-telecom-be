package service

import "github.com/greenblat17/stream-telecom/internal/repo"

type CampaignService struct {
	repos repo.Campaign
}

func NewCampaignService(repos *repo.Repository) *CampaignService {
	return &CampaignService{
		repos: repos,
	}
}
