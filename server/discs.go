package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"discgolfapi.com/m/database"
	"discgolfapi.com/m/models"
	"github.com/rs/zerolog/log"
)

// GetDiscs ... Get all discs
// @Summary Get all discs
// @Description Get all discs
// @Success 200 {array} models.Disc
// @Router /v1/discs [get]
func GetDiscs(w http.ResponseWriter, r *http.Request) {
	discs, err := database.Db.GetDiscs()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error when getting discs from database: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := models.DiscsResponse{Discs: discs}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error when marshalling json response: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// PutDiscs ... Add disc
// @Summary Add disc
// @Description Add disc
// @Accept json
// @Param disc body models.Disc true "Disc data"
// @Success 200 {object} object
// @Router /v1/discs [PUT]
func PutDiscs(w http.ResponseWriter, r *http.Request) {
	var disc models.Disc
	err := json.NewDecoder(r.Body).Decode(&disc)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error decoding disc from request: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.Db.PutDisc(&disc)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error when putting disc into database: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
