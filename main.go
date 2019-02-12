package main

import (
	"log"
	"net/http"

	"github.com/jeffveleze/gu_mvc/db"
)

func main() {

	dbClient := db.NewDbClient()

	defer dbClient.Database.Close()

	router := NewRouter(dbClient)

	log.Fatal(http.ListenAndServe(":8080", router))
}
