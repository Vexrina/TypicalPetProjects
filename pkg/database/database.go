package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func Connect(dbusr string, dbpwd string){
	conninfo := "user=root password=root host=127.0.0.1 sslmode=disable"
	
	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		log.Fatal(err)
	}

	dbName := "test_db"
	_, err = db.Exec("create database " +dbName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE example ( id integer, username varchar(255) )")

	if err != nil {
		log.Fatal(err)
	}
}