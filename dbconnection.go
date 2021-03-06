package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	dbDriver := "mysql"
	dbServer := "localhost:3307"
    dbUser := "root"
    dbPass := "system8"
    dbName := "go_db"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+ dbServer +")/"+dbName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}