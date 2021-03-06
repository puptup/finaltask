package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/puptup/finaltask/project/dbrepo"
)

var repDB dbrepo.DBWorker

func ConnectDBToHandlers(dbw dbrepo.DBWorker) {
	repDB = dbw
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func parseJsonToStruct(w http.ResponseWriter, r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}
	defer r.Body.Close()
	return nil
}
