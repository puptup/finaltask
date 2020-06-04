package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/FinalTask/dbrepo"
	"github.com/gorilla/mux"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups, err := dbrepo.GetGroups()
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get group")
		return
	}
	json.NewEncoder(w).Encode(groups)
}

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

	newGroup, err := dbrepo.PostGroup(group.Title)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to post group")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGroup)

}

func PutGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

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
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newGroup, err := dbrepo.PutGroup(id, group.Title)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update group")
		return
	}
	json.NewEncoder(w).Encode(newGroup)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = dbrepo.DeleteGroup(id)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
