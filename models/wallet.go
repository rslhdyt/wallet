package models

type Wallet struct {
	Id             string `json:"id"`
	PersonableId   string `json:"personable_id"`
	PersonableType string `json:"personable_type"`
	Balance        string `json:"balance"`
}
