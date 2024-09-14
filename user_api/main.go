package main

import (
    "log"
    "net/http"
    "user_api/db"
    "user_api/routes"
)

func main() {
    db.ConnectDB()
    router := routes.SetupRoutes()

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
