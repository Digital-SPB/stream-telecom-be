package service

import "github.com/greenblat17/stream-telecom/internal/repo"

type RegionService struct {
	repos repo.Regions
}

func NewRegionService(repos *repo.Repository) *RegionService {
	return &RegionService{
		repos: repos,
	}
}
