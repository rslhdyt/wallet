package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spenmo_wallet/database"
	"spenmo_wallet/models"

	"github.com/gorilla/mux"
)

func IndexWallet(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()
	selectQuery, err := db.Query("SELECT * FROM wallets ORDER BY id DESC")
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

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallets)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()

	requestBody, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	json.Unmarshal(requestBody, &wallet)

	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO wallets(personable_id, personable_type) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(wallet.PersonableId, wallet.PersonableType)
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wallet)
}

func ShowWallet(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()

	vars := mux.Vars(r)
	walletId := vars["id"]

	query, err := db.Query("SELECT * FROM wallets WHERE id=?", walletId)

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

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	db, _ := database.Connect()

	vars := mux.Vars(r)
	walletId := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	json.Unmarshal(requestBody, &wallet)

	if r.Method == "POST" {
		updateQuery, err := db.Prepare("UPDATE wallet SET personable_id=?, personable_type=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		updateQuery.Exec(wallet.PersonableId, wallet.PersonableType, walletId)
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wallet)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: users")
	json.NewEncoder(w).Encode(Users)
}
