package handler

import (
	"encoding/json"
	"net/http"
	"test-task_backend/events"
)

func GetCreatedLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(events.CollectionsCreated)
}

func GetMintedLog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(events.TokensMinted)
}
