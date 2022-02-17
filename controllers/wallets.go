package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"

	"github.com/gorilla/mux"
)

func IndexWallet(w http.ResponseWriter, r *http.Request) {
	selectQuery, err := database.Connector.Query("SELECT * FROM wallets ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	wallet := models.Wallet{}
	wallets := []models.Wallet{}

	for selectQuery.Next() {
		var id, personable_id, personable_type, balance string

		err = selectQuery.Scan(&id, &personable_id, &personable_type, &balance)

		if err != nil {
			panic(err.Error())
		}

		wallet.Id = id
		wallet.PersonableId = personable_id
		wallet.PersonableType = personable_type
		wallet.Balance = balance

		wallets = append(wallets, wallet)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallets)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	json.Unmarshal(requestBody, &wallet)

	if r.Method == "POST" {
		insForm, err := database.Connector.Prepare("INSERT INTO wallets(personable_id, personable_type) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(wallet.PersonableId, wallet.PersonableType)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wallet)
}

func ShowWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["id"]

	query, err := database.Connector.Query("SELECT * FROM wallets WHERE id=?", walletId)

	if err != nil {
		panic(err.Error())
	}

	wallet := models.Wallet{}

	for query.Next() {
		var id, personable_id, personable_type, balance string

		err = query.Scan(&id, &personable_id, &personable_type, &balance)

		if err != nil {
			panic(err.Error())
		}

		wallet.Id = id
		wallet.PersonableId = personable_id
		wallet.PersonableType = personable_type
		wallet.Balance = balance
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	json.Unmarshal(requestBody, &wallet)

	if r.Method == "PUT" {
		updateQuery, err := database.Connector.Prepare("UPDATE wallets SET personable_id=?, personable_type=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(wallet.PersonableId, wallet.PersonableType, walletId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["id"]

	if r.Method == "DELETE" {
		updateQuery, err := database.Connector.Prepare("DELETE FROM wallets WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(walletId)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
