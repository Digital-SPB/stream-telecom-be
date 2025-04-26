package repo

import (
	"github.com/greenblat17/stream-telecom/internal/model"
)

type Campaign interface {
	GetByID(id int64) (*model.Campaign, error)
	GetAllCampaigns() []*model.Campaign
}

type Click interface {
	GetClickDynamic(id int64) (*model.CampaignStats, error)
	GetByCampaignID(id int64) []*model.Click
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
