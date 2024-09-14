package db

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    createTable()
}

func createTable() {
    query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        age INT
    )`
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatal("Failed to create table:", err)
    }
}