package main

import (
	"log"
	"net/http"
	"user_api/db"
)

func main() {
	db.ConnectDB()

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
