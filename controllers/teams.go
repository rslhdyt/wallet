package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"

	"github.com/gorilla/mux"
)

func IndexTeam(w http.ResponseWriter, r *http.Request) {
	selectQuery, err := database.Connector.Query("SELECT * FROM teams ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	team := models.Team{}
	teams := []models.Team{}

	for selectQuery.Next() {
		var id, name, email string

		err = selectQuery.Scan(
			&id,
			&name,
			&email)

		if err != nil {
			panic(err.Error())
		}

		team.Id = id
		team.Name = name
		team.Email = email

		teams = append(teams, team)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var team models.Team
	json.Unmarshal(requestBody, &team)

	if r.Method == "POST" {
		insForm, err := database.Connector.Prepare("INSERT INTO teams(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(
			team.Name,
			team.Email)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func ShowTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["id"]

	query, err := database.Connector.Query("SELECT * FROM teams WHERE id=?", teamId)

	if err != nil {
		panic(err.Error())
	}

	team := models.Team{}

	for query.Next() {
		var id, name, email string

		err = query.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		team.Id = id
		team.Name = name
		team.Email = email
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var team models.Team
	json.Unmarshal(requestBody, &team)

	if r.Method == "PUT" {
		updateQuery, err := database.Connector.Prepare("UPDATE teams SET " +
			"name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(
			team.Name,
			team.Email,
			teamId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["id"]

	if r.Method == "DELETE" {
		updateQuery, err := database.Connector.Prepare("DELETE FROM teams WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(teamId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func TeamUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["id"]

	selectQuery, err := database.Connector.Query("SELECT * FROM users "+
		"LEFT JOIN users_teams ON users.id = users_teams.user_id WHERE users_teams.team_id = ?", teamId)

	if err != nil {
		panic(err.Error())
	}

	user := models.User{}
	users := []models.User{}

	for selectQuery.Next() {
		var id, name, email string

		err = selectQuery.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.Email = email

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
