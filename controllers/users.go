package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"

	"github.com/gorilla/mux"
)

func IndexUser(w http.ResponseWriter, r *http.Request) {
	selectQuery, err := database.Connector.Query("SELECT * FROM users ORDER BY id DESC")
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(requestBody, &user)

	if r.Method == "POST" {
		insForm, err := database.Connector.Prepare("INSERT INTO users(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(user.Name, user.Email)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	query, err := database.Connector.Query("SELECT * FROM users WHERE id=?", userId)

	if err != nil {
		panic(err.Error())
	}

	user := models.User{}

	for query.Next() {
		var id, name, email string

		err = query.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.Email = email
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var user models.User
	json.Unmarshal(requestBody, &user)

	if r.Method == "PUT" {
		updateQuery, err := database.Connector.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(user.Name, user.Email, userId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	if r.Method == "DELETE" {
		updateQuery, err := database.Connector.Prepare("DELETE FROM users WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(userId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func UserCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	cardQuery, err := database.Connector.Query("SELECT * FROM `cards` WHERE `wallet_id` IN "+
		"(SELECT id FROM `wallets` WHERE (`personable_id` IN "+
		"(SELECT team_id FROM `user_team` WHERE `user_id` = ?) "+
		"AND `personable_type` = 'Team') OR (`personable_id` = ? AND `personable_type` = 'User'))", userId, userId)

	if err != nil {
		panic(err.Error())
	}

	card := models.Card{}
	cards := []models.Card{}

	for cardQuery.Next() {
		var id, name, wallet_id, daily_limit, montlhy_limit string

		err = cardQuery.Scan(
			&id,
			&name,
			&wallet_id,
			&daily_limit,
			&montlhy_limit)

		if err != nil {
			panic(err.Error())
		}

		card.Id = id
		card.Name = name
		card.WalletId = wallet_id
		card.DailyLimit = daily_limit
		card.MonthlyLimit = montlhy_limit

		cards = append(cards, card)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}
