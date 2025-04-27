package model

type Region struct {
	ID   int64  `json:"region_id"`
	Name string `json:"name"`
}

type RegionMembersHeatMap struct {
	ID           int64  `json:"region_id"`
	Name         string `json:"name"`
	MembersCount int64  `json:"members_count"`
}

type CountClickByRegion struct {
	ID          int64  `json:"region_id"`
	Name        string `json:"name"`
	ClicksCount int64  `json:"clicks_count"`
}

type RegionInfo struct {
    ID          int    `json:"id"`
    NameEnglish string `json:"name_english"`
    UTCoffset   string `json:"utc_offset"`
}
