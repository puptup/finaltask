package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/FinalTask/dbrepo"
	"github.com/gorilla/mux"
)

func PostTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var timefr dbrepo.Timeframe
	err := parseJsonToStruct(w, r, &timefr)
	if err != nil {
		return
	}

	newTimeFrame, err := dbrepo.PostTimeFrame(timefr.TaskID, timefr.From, timefr.To)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to post timeframe")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTimeFrame)
}

func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = dbrepo.DeleteTimeFrame(id)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
