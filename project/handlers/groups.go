package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/puptup/FinalTask/project/dbrepo"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups, err := dbrepo.RepSQL.GetGroups()
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get group")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

func PostGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group dbrepo.Group
	err := parseJsonToStruct(w, r, &group)
	if err != nil {
		return
	}

	newGroup, err := dbrepo.RepSQL.PostGroup(group.Title)
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
	err := parseJsonToStruct(w, r, &group)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newGroup, err := dbrepo.RepSQL.PutGroup(id, group.Title)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update group")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newGroup)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = dbrepo.RepSQL.DeleteGroup(id)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
