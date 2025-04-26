package repo

import "context"

type Campaign interface {
	GetCampainActivity(id int, ctx context.Context)
}

type Click interface {
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
