package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "user_api/db"
    "user_api/models"
    "github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Insert user into database
    query := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id"
    err = db.DB.QueryRow(query, user.Name, user.Email, user.Age).Scan(&user.ID)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    query := "SELECT id, name, email, age FROM users WHERE id=$1"
    row := db.DB.QueryRow(query, userID)
    err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    query := "UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4"
    _, err = db.DB.Exec(query, user.Name, user.Email, user.Age, userID)
    if err != nil {
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    query := "DELETE FROM users WHERE id=$1"
    _, err = db.DB.Exec(query, userID)
    if err != nil {
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
