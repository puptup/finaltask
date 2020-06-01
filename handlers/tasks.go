package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/FinalTask/dbrepo"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := dbrepo.GetTasks()

	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task dbrepo.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if task.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if task.GroupID == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newTask := dbrepo.PostTask(task.Title, task.GroupID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var task dbrepo.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if task.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if task.GroupID == 0 || id == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newTask := dbrepo.PutTask(id, task.GroupID, task.Title)
	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
