package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/FinalTask/dbrepo"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := dbrepo.GetTasks()
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get tasks")
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task dbrepo.Task
	err := parseJsonToStruct(w, r, &task)
	if err != nil {
		return
	}

	newTask, err := dbrepo.PostTask(task.Title, task.GroupID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to post task")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var task dbrepo.Task
	err := parseJsonToStruct(w, r, &task)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newTask, err := dbrepo.PutTask(id, task.GroupID, task.Title)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update task")
		return
	}
	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = dbrepo.DeleteTask(id)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Failed to delete")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
