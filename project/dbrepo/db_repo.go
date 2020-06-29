package dbrepo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var RepSQL DBWorker = &DBRepo{}

//DBRepo structure for interacting with the database
type DBRepo struct {
	DB *sql.DB
}

//DBInit открытие соединиение с БД. В ней же и устанавливается DBSet.
func (repo *DBRepo) DBInit() *sql.DB {
	config := DBConfig{}

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetDBName(), config.GetPort())
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
	}

	repo.DB = db
	return db
}
