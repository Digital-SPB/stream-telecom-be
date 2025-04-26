package repo

import (
<<<<<<< HEAD
	"context"

=======
>>>>>>> cfb156695d7792d1a62c696ed040c9b8e24d943c
	"github.com/greenblat17/stream-telecom/internal/model"
)

type Campaign interface {
	GetByID(id int64) (*model.Campaign, error)
	GetAllCampaigns() []*model.Campaign
}

type Click interface {
<<<<<<< HEAD
	GetClickDynamic(id int64) (*model.CampaignStats, error)
=======
	GetByCampaignID(id int64) []*model.Click
>>>>>>> cfb156695d7792d1a62c696ed040c9b8e24d943c
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
