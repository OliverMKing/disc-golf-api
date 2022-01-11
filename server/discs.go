package server

import (
	"encoding/json"
	"net/http"

	. "discgolfapi.com/m/models"
)

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
