package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/puptup/FinalTask/project/dbrepo"
	"github.com/puptup/FinalTask/project/handlers"
)

func main() {
	connection := dbrepo.RepSQL.DBInit()
	defer connection.Close()

	router := mux.NewRouter()
	GroupsRouter := router.PathPrefix("/groups").Subrouter()
	TasksRouter := router.PathPrefix("/tasks").Subrouter()
	TimeframesRouter := router.PathPrefix("/timeframes").Subrouter()

	GroupsRouter.HandleFunc("", handlers.GetGroups).Methods(http.MethodGet)
	GroupsRouter.HandleFunc("/", handlers.PostGroup).Methods(http.MethodPost)
	GroupsRouter.HandleFunc("/{id}", handlers.PutGroup).Methods(http.MethodPut)
	GroupsRouter.HandleFunc("/{id}", handlers.DeleteGroup).Methods(http.MethodDelete)

	TasksRouter.HandleFunc("", handlers.GetTasks).Methods(http.MethodGet)
	TasksRouter.HandleFunc("/", handlers.PostTask).Methods(http.MethodPost)
	TasksRouter.HandleFunc("/{id}", handlers.PutTask).Methods(http.MethodPut)
	TasksRouter.HandleFunc("/{id}", handlers.DeleteTask).Methods(http.MethodDelete)

	TimeframesRouter.HandleFunc("/", handlers.PostTimeframe).Methods(http.MethodPost)
	TimeframesRouter.HandleFunc("/{id}", handlers.DeleteTimeframe).Methods(http.MethodDelete)

	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
