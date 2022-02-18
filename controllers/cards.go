package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"

	"github.com/gorilla/mux"
)

func IndexCard(w http.ResponseWriter, r *http.Request) {
	selectQuery, err := database.Connector.Query("SELECT * FROM cards ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	card := models.Card{}
	cards := []models.Card{}

	for selectQuery.Next() {
		var id, wallet_id int64
		var name string
		var daily_limit, monthly_limit float64

		err = selectQuery.Scan(
			&id,
			&name,
			&wallet_id,
			&daily_limit,
			&monthly_limit)

		if err != nil {
			panic(err.Error())
		}

		card.Id = id
		card.Name = name
		card.WalletId = wallet_id
		card.DailyLimit = daily_limit
		card.MonthlyLimit = monthly_limit

		cards = append(cards, card)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cards)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var card models.Card
	json.Unmarshal(requestBody, &card)

	if r.Method == "POST" {
		insForm, err := database.Connector.Prepare("INSERT INTO cards(name, wallet_id, daily_limit, monthly_limit) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		result, err := insForm.Exec(
			card.Name,
			card.WalletId,
			card.DailyLimit,
			card.MonthlyLimit)

		if err != nil {
			panic(err.Error())
		}

		card.Id, _ = result.LastInsertId()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

func ShowCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["id"]

	query, err := database.Connector.Query("SELECT * FROM cards WHERE id=?", cardId)

	if err != nil {
		panic(err.Error())
	}

	card := models.Card{}

	for query.Next() {
		var id, wallet_id int64
		var name string
		var daily_limit, monthly_limit float64

		err = query.Scan(&id, &name, &wallet_id, &daily_limit, &monthly_limit)

		if err != nil {
			panic(err.Error())
		}

		card.Id = id
		card.Name = name
		card.WalletId = wallet_id
		card.DailyLimit = daily_limit
		card.MonthlyLimit = monthly_limit
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var card models.Card
	json.Unmarshal(requestBody, &card)

	if r.Method == "PUT" {
		updateQuery, err := database.Connector.Prepare("UPDATE cards SET " +
			"name=?, wallet_id=?, daily_limit=?, monthly_limit=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(
			card.Name,
			card.WalletId,
			card.DailyLimit,
			card.MonthlyLimit,
			cardId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(card)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cardId := vars["id"]

	if r.Method == "DELETE" {
		updateQuery, err := database.Connector.Prepare("DELETE FROM cards WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(cardId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
