package main

import (
	"fmt"
	"log"
	"net/http"

	"spenmo_wallet/controllers"

	"github.com/gorilla/mux"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Card struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	WalletId     string `json:"wallet_id"`
	DailyLimit   string `json:"daily_limit"`
	MonthlyLimit string `json:"monthly_id"`
}

var Users []User
var Cards []Card

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: rootPage")
}

// func userShowPage(w http.ResponseWriter, r *http.Request) {
// 	userID := mux.Vars(r)["id"]

// 	for _, user := range Users {
// 		if user.Id == userID {
// 			json.NewEncoder(w).Encode(user)
// 		}
// 	}
// }

// func userCardsPage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: users")
// 	json.NewEncoder(w).Encode(Users)
// }

// func userCardsCreatePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: users")
// 	json.NewEncoder(w).Encode(Users)
// }

func main() {
	Users = []User{
		{Id: "1", Name: "John", Email: "banana@mail.com"},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootPage)

	// router.HandleFunc("/users", usersPage).Methods("POST")
	// router.HandleFunc("/users/{id}", usersPage).Methods("GET")

	router.HandleFunc("/wallets", controllers.IndexWallet).Methods("GET")
	router.HandleFunc("/wallets", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{id}", controllers.ShowWallet).Methods("GET")
	// router.HandleFunc("/wallets/{id}", updateWalletsApi).Methods("PUT")
	// router.HandleFunc("/wallets/{id}", deleteWalletsApi).Methods("DELETE")

	// router.HandleFunc("/cards", usersPage)
	// router.HandleFunc("/cards", userCardsPage).Methods("POST")
	// router.HandleFunc("/cards/{id}", userShowPage).Methods("GET")
	// router.HandleFunc("/cards/{id}", userCardsCreatePage).Methods("PUT")
	// router.HandleFunc("/cards/{id}", userCardsCreatePage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
