package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nanoohlaing1997/order-api/api"
	"github.com/nanoohlaing1997/order-api/database"
)

func main() {
	fmt.Println("Ordering API application is ready to serve")
	godotenv.Load()

	router := RegisterRoute()

	fmt.Printf("Starting server on port %s", os.Getenv("REST_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("REST_PORT")), router)
}

func RegisterRoute() *mux.Router {
	dbm := database.NewDatabaseManager()
	controller := api.NewControllerManager(dbm)

	router := mux.NewRouter()
	// Registering 3 APIs
	router.HandleFunc("/orders", controller.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", controller.TakeOrder).Methods("PATCH")
	router.HandleFunc("/orders", controller.ListOrder).
		Queries("page", "{page}", "limit", "{limit}").
		Methods("GET")

	return router
}
