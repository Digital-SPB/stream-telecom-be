package model

type Region struct {
	ID int64 `json:"region_id"`
	Name string	`json:"name"`
}

type RegionHeatMap struct {
	ID int64 `json:"region_id"`
	Name string `json:"name"`
	ClickCount int64 `json:"click_count"`
}
