package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var repo *sql.DB

func DBSet(db *sql.DB) {
	repo = db
}

func DBInit() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	DBSet(db)
	return db
}
