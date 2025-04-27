package model

import (
	"time"
)

type Click struct {
	ID         string    `json:"uid"`
	ClickDate  time.Time `json:"click_date"`
	ClickTime  time.Time `json:"click_time"`
	MemberID   int64     `json:"member_id"`
	CampaignID int64     `json:"campaign_id"`
	RegionID   int64     `json:"region"`
	OS         string    `json:"os"`
	Browser    string    `json:"browser"`
	UserAgent  string    `json:"user_agent"`
	Language   string    `json:"language"`
	Device     string    `json:"device"`
}

// три структуры для 2 ручки
type DailyStat struct {
	Date        time.Time `json:"date"`
	ClicksCount int       `json:"clicks_count"`
	Percentage  float64   `json:"percentage"`
}

type MonthlyStat struct {
	Month       time.Time `json:"month"`
	ClicksCount int       `json:"clicks_count"`
	Percentage  float64   `json:"percentage"`
}

type CampaignStats struct {
	DailyStats   []*DailyStat   `json:"daily_stats"`
	MonthlyStats []*MonthlyStat `json:"monthly_stats"`
	TotalClicks  int            `json:"total_clicks"`
}
