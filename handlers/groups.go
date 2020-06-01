package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FinalTask/dbrepo"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups := dbrepo.GetGroups()

	json.NewEncoder(w).Encode(groups)
}

/*
{
	“title”: “group’s title”
}
*/

func PostGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group dbrepo.Group
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&group); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if group.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newGroup := dbrepo.PostGroup(group.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGroup)

}

func PutGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
