package repo

import (
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type Campaign interface {
	GetByID(id int64) (*model.Campaign, error)
	GetAllCampaigns() []*model.Campaign
	GetCreationDynamic(start time.Time, end time.Time, intervalType string) ([]*model.IntervalResult, error)
}

type Click interface {
	GetClickDynamic(id int64) (*model.CampaignStats, error)
	GetByCampaignID(id int64) []*model.Click
	GetAll() []*model.Click
	GetTimeActivity() *model.TimeActivityResponse
	GetDailyTimeActivity(targetDate time.Time) *model.DailyTimeActivityResponse
}

type Regions interface {
	GetAll() []*model.Region
	GetRegionsInfo() []*model.RegionInfo 
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
