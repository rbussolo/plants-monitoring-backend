package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "bussolo"
	dbname   = "weather"
)

var instance *sqlx.DB = nil

func GetInstance() *sqlx.DB {
	if instance == nil {
		instance = connectDataBase()
	}

	return instance
}

func connectDataBase() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
