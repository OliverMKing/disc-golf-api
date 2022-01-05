package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "discgolfapi.com/m/models"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/discs", Discs).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is running and healthy")
}

func Discs(w http.ResponseWriter, r *http.Request) {
	teebird := Disc{Name: "Teebird"}
	buzzz := Disc{Name: "Buzzz"}
	zone := Disc{Name: "Zone"}
	discs := []Disc{teebird, buzzz, zone}

	resp := DiscResponse{Discs: discs}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
