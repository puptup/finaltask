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

func PostGroup(w http.ResponseWriter, r *http.Request) {

}

func PutGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
