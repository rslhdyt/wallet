package models

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cards []Card `json:"cards"`
	Teams []Team `json:"teams"`
}
