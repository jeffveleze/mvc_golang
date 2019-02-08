package main

import (
	"log"
	"net/http"

	"github.com/jeffveleze/gu_mvc/db"
)

func main() {

	dbClient := db.NewDbClient()

	router := NewRouter(dbClient)

	// defer dbClient.Database.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
