package model

import (
	"time"
)

type Click struct {
	ID string `json:"uid"`
	ClickDate time.Time `json:"click_date"`
	ClickTime time.Time `json:"click_time"`
	MemberID int64 `json:"member_id"`
	CampaignID int64 `json:"campaign_id"`
	RegionID int64 `json:"region"`
	OS string `json:"os"`
	Browser string `json:"browser"`
	UserAgent string `json:"user_agent"`
	Language string `json:"language"`
	Device string `json:"device"`
}