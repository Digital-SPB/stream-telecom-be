package repo

import (
	"context"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type Campaign interface {
	GetCampainActivity(id int, ctx context.Context)
}

type Click interface {
	GetClickDynamic(id int64) (*model.CampaignStats, error)
}

type Regions interface {
}

type Repository struct {
	Campaign
	Click
	Regions
}

func NewRepository() *Repository {
	return &Repository{
		Campaign: LoadCampaignRepo(),
		Click:    LoadClickRepo(),
		Regions:  LoadRegionsRepo(),
	}
}
