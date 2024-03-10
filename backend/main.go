package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3307)/books")

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Print("Connected to the database.")
}
