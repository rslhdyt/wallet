package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"
)

func CreateUserTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var userTeam models.UserTeam
	json.Unmarshal(requestBody, &userTeam)

	if r.Method == "POST" {
		insForm, err := database.Connector.Prepare("INSERT INTO user_team(user_id, team_id) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(userTeam.UserId, userTeam.TeamId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userTeam)
}

func DeleteUserTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var userTeam models.UserTeam
	json.Unmarshal(requestBody, &userTeam)

	if r.Method == "DELETE" {
		updateQuery, err := database.Connector.Prepare("DELETE FROM user_team WHERE user_id=? AND team_id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(userTeam.UserId, userTeam.TeamId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
