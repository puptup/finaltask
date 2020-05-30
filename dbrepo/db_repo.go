package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DBRepo struct {
	DB *sql.DB
}

var repo = DBRepo{}

func DBInit() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	repo.DB = db

	return db
}
