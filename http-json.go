package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	type Health struct {
		Alive bool `json:"alive"`
	}

	status := Health{Alive: true}
	j, _ := json.Marshal(status)

	w.Write(j)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
