package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
)

type ActivityService struct {
	campaigns map[int64]*model.Campaign
	clicks    []*model.Click
}

type ActivityMetrics struct {
	CampaignID    int64     `json:"campaign_id"`
	CreatedAt     time.Time `json:"created_at"`
	TotalClicks   int       `json:"total_clicks"`
	HourlyClicks  []int     `json:"hourly_clicks"`
	TimeRange     []string  `json:"time_range"`
}

func NewActivityService() *ActivityService {
	return &ActivityService{
		campaigns: make(map[int64]*model.Campaign),
		clicks:    make([]*model.Click, 0),
	}
}

func (s *ActivityService) LoadData() error {
	// Load campaigns
	if err := s.loadCampaigns(); err != nil {
		return err
	}

	// Load clicks
	if err := s.loadClicks(); err != nil {
		return err
	}

	return nil
}

func (s *ActivityService) loadCampaigns() error {
	file, err := os.Open("data/campaign.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			continue
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			continue
		}

		s.campaigns[id] = &model.Campaign{
			ID:        id,
			CreatedAt: createdAt,
		}
	}

	return nil
}

func (s *ActivityService) loadClicks() error {
	file, err := os.Open("data/clicks.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}


		clickDate, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			log.Println("err", err)
			continue
		}

		clickTime, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			log.Println("err", err)
			continue
		}

		uid := record[2]


		memberID, err := strconv.ParseInt(record[3], 10, 64)
		if err != nil {
			log.Println("err", err)
			continue
		}

		campaignID, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			log.Println("err", err)
			continue
		}

		regionID, err := strconv.ParseInt(record[5], 10, 64)
		if err != nil {
			log.Println("err", err)
			continue
		}

		click := &model.Click{
			ID:         uid,
			ClickDate:  clickDate,
			ClickTime:  clickTime,
			MemberID:   memberID,
			CampaignID: campaignID,
			RegionID:   regionID,
			OS:         record[6],
			Browser:    record[7],
			UserAgent:  record[8],
			Language:   record[9],
			Device:     record[10],
		}

		s.clicks = append(s.clicks, click)
	}

	return nil
}

func (s *ActivityService) GetCampaignActivity(campaignID int64, countHours int64) (*ActivityMetrics, error) {
	campaign, exists := s.campaigns[campaignID]
	if !exists {
		return nil, nil
	}

	metrics := &ActivityMetrics{
		CampaignID:   campaignID,
		CreatedAt:    campaign.CreatedAt,
		HourlyClicks: make([]int, countHours),
		TimeRange:    make([]string, countHours),
	}

	// Calculate time range for hours
	startTime := campaign.CreatedAt
	
	for i := 0; i < int(countHours); i++ {
		metrics.TimeRange[i] = startTime.Add(time.Duration(i) * time.Hour).Format("15:04")
	}

	// Count clicks within specified hours
	for _, click := range s.clicks {
		if click.CampaignID != campaignID {
			continue
		}

		clickDateTime := time.Date(
			click.ClickDate.Year(),
			click.ClickDate.Month(),
			click.ClickDate.Day(),
			click.ClickTime.Hour(),
			click.ClickTime.Minute(),
			click.ClickTime.Second(),
			0,
			time.UTC,
		)

		if clickDateTime.Before(startTime) || clickDateTime.After(startTime.Add(time.Duration(countHours) * time.Hour)) {
			continue
		}

		metrics.TotalClicks++
		hourIndex := int(clickDateTime.Sub(startTime).Hours())
		if hourIndex >= 0 && hourIndex < int(countHours) {
			metrics.HourlyClicks[hourIndex]++
		}
	}

	return metrics, nil
}

func (s *ActivityService) GetAllCampaignsActivity() ([]*ActivityMetrics, error) {
	var metrics []*ActivityMetrics

	for _, campaign := range s.campaigns {
		campaignMetrics, err := s.GetCampaignActivity(campaign.ID, 4)
		if err != nil {
			return nil, err
		}
		if campaignMetrics != nil {
			metrics = append(metrics, campaignMetrics)
		}
	}

	return metrics, nil
} 