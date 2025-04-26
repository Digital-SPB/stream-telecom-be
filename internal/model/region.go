package model

type Region struct {
	ID int64 `json:"region_id"`
	Name string	`json:"name"`
}

type RegionMembersHeatMap struct {
	ID int64 `json:"region_id"`
	Name string `json:"name"`
	MembersCount int64 `json:"members_count"`
}
