package models

type Card struct {
	Id           string `json:"id"`
	WalletId     string `json:"wallet_id"`
	Name         string `json:"name"`
	DailyLimit   string `json:"daily_limit"`
	MonthlyLimit string `json:"montlhy_limit"`
}
