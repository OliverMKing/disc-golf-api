package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"discgolfapi.com/m/models"
	"github.com/rs/zerolog/log"
)

// Discs ... Get all discs
// @Summary Get all discs
// @Description Get all discs
// @Success 200 {array} models.Disc
// @Router /discs [get]
func Discs(w http.ResponseWriter, r *http.Request) {
	discs, err := Db.GetDiscs()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error when getting discs from database: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := models.DiscsResponse{Discs: discs}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error when marshalling json response: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
