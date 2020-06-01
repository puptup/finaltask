package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FinalTask/dbrepo"
)

func PostTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var timefr dbrepo.Timeframe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&timefr); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if timefr.From == "" || timefr.To == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if timefr.TaskID == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newTimeFrame := dbrepo.PostTimeFrame(timefr.TaskID, timefr.From, timefr.To)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTimeFrame)
}

func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {

}
