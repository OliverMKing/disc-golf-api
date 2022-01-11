package server

import (
	"encoding/json"
	"net/http"
)

// Discs ... Get all discs
// @Summary Get all discs
// @Description Get all discs
// @Success 200 {array} models.Disc
// @Router /discs [get]
func Discs(w http.ResponseWriter, r *http.Request) {
	resp, err := Db.GetDiscs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
