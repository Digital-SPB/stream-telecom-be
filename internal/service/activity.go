package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

type FormattedDuration struct {
	Years   int `json:"years,omitempty"`
	Months  int `json:"months,omitempty"`
	Days    int `json:"days,omitempty"`
	Hours   int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`
	Seconds int `json:"seconds,omitempty"`
}

type ReactionTimeMetrics struct {
	CampaignID      int64             `json:"campaign_id"`
	AverageTime     FormattedDuration `json:"average_time"`
	MedianTime      FormattedDuration `json:"median_time"`
	MinTime         FormattedDuration `json:"min_time"`
	MaxTime         FormattedDuration `json:"max_time"`
	Percentile90th  FormattedDuration `json:"percentile_90th"`
	Percentile95th  FormattedDuration `json:"percentile_95th"`
	TotalCustomers  int               `json:"total_customers"`
}

type NoClicksFoundError struct {
	CampaignID int64
	Message    string
}

type CampaignList struct {
	Total       int            `json:"total"`
	Page        int            `json:"page"`
	PerPage     int            `json:"per_page"`
	TotalPages  int            `json:"total_pages"`
	Campaigns   []CampaignInfo `json:"campaigns"`
}

type CampaignInfo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func (e *NoClicksFoundError) Error() string {
	return fmt.Sprintf("campaign %d: %s", e.CampaignID, e.Message)
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
			Name:      GenerateCampaignName(id),
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

func formatDuration(seconds float64) FormattedDuration {
	// Convert to integer seconds
	totalSeconds := int(seconds)
	
	// Calculate each unit
	secondsInMinute := 60
	secondsInHour := 60 * secondsInMinute
	secondsInDay := 24 * secondsInHour
	secondsInMonth := 30 * secondsInDay    // Approximation
	secondsInYear := 365 * secondsInDay    // Approximation

	years := totalSeconds / secondsInYear
	totalSeconds = totalSeconds % secondsInYear

	months := totalSeconds / secondsInMonth
	totalSeconds = totalSeconds % secondsInMonth

	days := totalSeconds / secondsInDay
	totalSeconds = totalSeconds % secondsInDay

	hours := totalSeconds / secondsInHour
	totalSeconds = totalSeconds % secondsInHour

	minutes := totalSeconds / secondsInMinute
	remainingSeconds := totalSeconds % secondsInMinute

	return FormattedDuration{
		Years:   years,
		Months:  months,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: remainingSeconds,
	}
}

func (s *ActivityService) GetCustomerReactionTime(campaignID int64) (*ReactionTimeMetrics, error) {
	campaign, exists := s.campaigns[campaignID]
	if !exists {
		return nil, &NoClicksFoundError{
			CampaignID: campaignID,
			Message:    "campaign not found",
		}
	}

	// Map to store first click time for each member
	memberFirstClicks := make(map[int64]time.Duration)

	// Find first click for each member
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

		// Skip clicks that happened before campaign creation
		if clickDateTime.Before(campaign.CreatedAt) {
			continue
		}

		reactionTime := clickDateTime.Sub(campaign.CreatedAt)
		
		// If this is the first click for this member or earlier than previous first click
		if currentTime, exists := memberFirstClicks[click.MemberID]; !exists || reactionTime < currentTime {
			memberFirstClicks[click.MemberID] = reactionTime
		}
	}

	// Convert map to slice for statistical calculations
	var reactionTimes []float64
	for _, duration := range memberFirstClicks {
		reactionTimes = append(reactionTimes, duration.Seconds())
	}

	// If no valid clicks found
	if len(reactionTimes) == 0 {
		return nil, &NoClicksFoundError{
			CampaignID: campaignID,
			Message:    "no valid clicks found after campaign creation date",
		}
	}

	// Sort for percentile calculations
	sort.Float64s(reactionTimes)

	// Calculate statistics
	var sum float64
	for _, t := range reactionTimes {
		sum += t
	}

	metrics := &ReactionTimeMetrics{
		CampaignID:     campaignID,
		AverageTime:    formatDuration(sum / float64(len(reactionTimes))),
		MedianTime:     formatDuration(calculateMedian(reactionTimes)),
		MinTime:        formatDuration(reactionTimes[0]),
		MaxTime:        formatDuration(reactionTimes[len(reactionTimes)-1]),
		Percentile90th: formatDuration(calculatePercentile(reactionTimes, 90)),
		Percentile95th: formatDuration(calculatePercentile(reactionTimes, 95)),
		TotalCustomers: len(reactionTimes),
	}

	return metrics, nil
}

func calculateMedian(sorted []float64) float64 {
	length := len(sorted)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		return (sorted[length/2-1] + sorted[length/2]) / 2
	}
	return sorted[length/2]
}

func calculatePercentile(sorted []float64, percentile float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	index := (percentile / 100) * float64(len(sorted)-1)
	i := int(math.Floor(index))
	if i+1 >= len(sorted) {
		return sorted[len(sorted)-1]
	}
	fraction := index - float64(i)
	return sorted[i] + fraction*(sorted[i+1]-sorted[i])
}

func (s *ActivityService) GetAllCampaigns(page, perPage int) *CampaignList {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10 // значение по умолчанию
	}

	campaigns := make([]CampaignInfo, 0, len(s.campaigns))
	
	// Собираем все кампании
	for id, campaign := range s.campaigns {
		campaigns = append(campaigns, CampaignInfo{
			ID:        id,
			Name:      campaign.Name,
			CreatedAt: campaign.CreatedAt,
		})

	}

	// Сортируем по ID для консистентности
	sort.Slice(campaigns, func(i, j int) bool {
		return campaigns[i].ID < campaigns[j].ID
	})

	// Вычисляем общее количество страниц
	total := len(campaigns)
	totalPages := (total + perPage - 1) / perPage

	// Вычисляем начальный и конечный индексы для текущей страницы
	start := (page - 1) * perPage
	end := start + perPage
	if end > total {
		end = total
	}
	if start > total {
		start = total
	}

	return &CampaignList{
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		Campaigns:  campaigns[start:end],
	}
} 