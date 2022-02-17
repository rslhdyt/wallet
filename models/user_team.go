package models

type UserTeam struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	TeamId string `json:"team_id"`
}
