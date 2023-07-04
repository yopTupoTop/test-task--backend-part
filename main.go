package main

import (
	"net/http"
	"test-task_backend/events"
	"test-task_backend/handler"

	"github.com/gorilla/mux"
)

func main() {
	go events.SubscribeToEvents()
	router := mux.NewRouter()
	router.HandleFunc("/collectionsCreated", handler.GetCreatedLog).Methods("GET")
	router.HandleFunc("/tokensMinted", handler.GetMintedLog).Methods("GET")

	http.ListenAndServe("localhost:8080", router)
}
