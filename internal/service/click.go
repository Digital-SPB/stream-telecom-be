package service

import (
	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

type ClickService struct {
	repos repo.Click
}

func NewClickService(repos *repo.Repository) *ClickService {
	return &ClickService{
		repos: repos,
	}
}

func (s *ClickService) GetClickDynamic(id int64) (*model.CampaignStats, error) {
	return s.repos.GetClickDynamic(id)
}
