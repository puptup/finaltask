package handlers

import (
	"fmt"
	"net/http"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "hi there")
}

func PostGroup(w http.ResponseWriter, r *http.Request) {

}

func PutGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}
