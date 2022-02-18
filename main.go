package main

import (
	"fmt"
	"log"
	"net/http"

	"spenmo_wallet/controllers"
	"spenmo_wallet/database"

	"github.com/gorilla/mux"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: rootPage")
}

func main() {
	// init DB
	database.Connect()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootPage)

	router.HandleFunc("/users", controllers.IndexUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.ShowUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/cards", controllers.UserCards).Methods("GET")

	router.HandleFunc("/teams", controllers.IndexTeam).Methods("GET")
	router.HandleFunc("/teams", controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/teams/{id}", controllers.ShowTeam).Methods("GET")
	router.HandleFunc("/teams/{id}", controllers.UpdateTeam).Methods("PUT")
	router.HandleFunc("/teams/{id}", controllers.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/teams/{id}/users", controllers.TeamUsers).Methods("GET")

	router.HandleFunc("/user_teams", controllers.CreateUserTeam).Methods("POST")
	router.HandleFunc("/user_teams", controllers.DeleteUserTeam).Methods("DELETE")

	router.HandleFunc("/wallets", controllers.IndexWallet).Methods("GET")
	router.HandleFunc("/wallets", controllers.CreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{id}", controllers.ShowWallet).Methods("GET")
	router.HandleFunc("/wallets/{id}", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/wallets/{id}", controllers.DeleteWallet).Methods("DELETE")

	router.HandleFunc("/cards", controllers.IndexCard).Methods("GET")
	router.HandleFunc("/cards", controllers.CreateCard).Methods("POST")
	router.HandleFunc("/cards/{id}", controllers.ShowCard).Methods("GET")
	router.HandleFunc("/cards/{id}", controllers.UpdateCard).Methods("PUT")
	router.HandleFunc("/cards/{id}", controllers.DeleteCard).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", Limit(router)))
}
