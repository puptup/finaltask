package dbrepo

import "os"

type DBConfig struct {
}

func (db *DBConfig) GetUser() string {
	return os.Getenv("POSTGRES_USER")
}

func (db *DBConfig) GetHost() string {
	return os.Getenv("POSTGRES_HOST")
}

func (db *DBConfig) GetPort() string {
	return os.Getenv("POSTGRES_PORT")
}

func (db *DBConfig) GetPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

func (db *DBConfig) GetDBName() string {
	return os.Getenv("POSTGRES_DB")
}
