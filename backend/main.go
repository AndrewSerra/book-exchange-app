package main

import (
	"database/sql"
	"log"

	"./api"
	"./db"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func testDbConn(db *sql.DB) {
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Print("Connected to the database.")
}

func main() {

	var err error
	dbConn, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3307)/books")

	if err != nil {
		log.Fatal(err)
	}

	testDbConn(dbConn)

	router := gin.Default()

	apiController := &api.APIContoller{
		DBController: &db.DBController{Database: dbConn},
	}

	router.GET("/users/:id", apiController.GetUserByID)
	router.POST("/users", apiController.CreateUser)

	router.Run()
}
