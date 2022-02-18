package models

type Card struct {
	Id           int64   `json:"id"`
	WalletId     int64   `json:"wallet_id"`
	Name         string  `json:"name"`
	DailyLimit   float64 `json:"daily_limit"`
	MonthlyLimit float64 `json:"montlhy_limit"`
}
