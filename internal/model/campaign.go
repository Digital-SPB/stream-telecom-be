package model

import "time"

type Campaign struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
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