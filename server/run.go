package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "discgolfapi.com/m/docs"
	. "discgolfapi.com/m/models"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/discs", Discs).Methods("GET")

	router.PathPrefix("/swagger/").Handler((httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	)))

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

// HeathCheck ... Get status of server
// @Summary Get status of server
// @Description Get API health status
// @Success 200
// @Router /health-check [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is running and healthy")
}

// Discs ... Get all discs
// @Summary Get all discs
// @Description Get all discs
// @Success 200 {array} models.Disc
// @Router /discs [get]
func Discs(w http.ResponseWriter, r *http.Request) {
	teebird := Disc{Name: "Teebird"}
	buzzz := Disc{Name: "Buzzz"}
	zone := Disc{Name: "Zone"}
	discs := []Disc{teebird, buzzz, zone}

	resp := DiscsResponse{Discs: discs}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
