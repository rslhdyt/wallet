package models

type Wallet struct {
	Id             int64   `json:"id"`
	PersonableId   int64   `json:"personable_id"`
	PersonableType string  `json:"personable_type"`
	Balance        float64 `json:"balance"`
}
