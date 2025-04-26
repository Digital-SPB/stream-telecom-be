package model

import "time"

type Campaign struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
