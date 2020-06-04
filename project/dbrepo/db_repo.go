package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var repo *sql.DB

//DBSet установит объект *DB, через который можно будет взаимодействовать с базой данных
func DBSet(db *sql.DB) {
	repo = db
}

//DBInit открытие соединиение с БД. В ней же и устанавливается DBSet.
func DBInit() *sql.DB {
	config := DBConfig{}

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetDBName(), config.GetPort())
	log.Println(dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	DBSet(db)
	return db
}
