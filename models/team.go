package models

type Team struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}
