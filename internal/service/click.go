package service

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/greenblat17/stream-telecom/internal/model"
	"github.com/greenblat17/stream-telecom/internal/repo"
)

type ClickService struct {
	clickRepository repo.Click
	campaignRepository repo.Campaign
}

func NewClickService(repos *repo.Repository) *ClickService {
	return &ClickService{
		clickRepository: repos.Click,
		campaignRepository: repos.Campaign,
	}
}

func formatDuration(seconds float64) model.FormattedDuration {
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

	return model.FormattedDuration{
		Years:   years,
		Months:  months,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: remainingSeconds,
	}
}

func (s *ClickService) GetCustomerReactionTime(campaignID int64) (*model.ReactionTimeMetrics, error) {
	campaign, err := s.campaignRepository.GetByID(campaignID)
	if err != nil {
		return nil, fmt.Errorf("campaign not found")
	}

	// Map to store first click time for each member
	memberFirstClicks := make(map[int64]time.Duration)

	// Find first click for each member
	clicks := s.clickRepository.GetByCampaignID(campaignID)
	for _, click := range clicks {

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
		return nil, fmt.Errorf("no valid clicks found after campaign creation date")
	}

	// Sort for percentile calculations
	sort.Float64s(reactionTimes)

	// Calculate statistics
	var sum float64
	for _, t := range reactionTimes {
		sum += t
	}

	metrics := &model.ReactionTimeMetrics{
		CampaignID:     campaignID,
		AverageTime:    formatDuration(sum / float64(len(reactionTimes))),
		MedianTime:     formatDuration(calculateClickMedian(reactionTimes)),
		MinTime:        formatDuration(reactionTimes[0]),
		MaxTime:        formatDuration(reactionTimes[len(reactionTimes)-1]),
		Percentile90th: formatDuration(calculateClickPercentile(reactionTimes, 90)),
		Percentile95th: formatDuration(calculateClickPercentile(reactionTimes, 95)),
		TotalCustomers: len(reactionTimes),
	}

	return metrics, nil
}

// calculateMedian returns the median value from a sorted slice of float64 values
func calculateClickMedian(sorted []float64) float64 {
	length := len(sorted)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		return (sorted[length/2-1] + sorted[length/2]) / 2
	}
	return sorted[length/2]
}

// calculatePercentile returns the percentile value from a sorted slice of float64 values
func calculateClickPercentile(sorted []float64, percentile float64) float64 {
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