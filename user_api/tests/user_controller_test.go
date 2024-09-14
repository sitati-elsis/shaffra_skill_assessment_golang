package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"user_management_api/db"
	"user_management_api/models"
	"user_management_api/routes"
)

var router *http.ServeMux

// TestMain sets up the environment before running tests
func TestMain(m *testing.M) {
	// Connect to the test database
	db.ConnectDB()

	// Set up routes
	router = routes.SetupRoutes()

	// Run tests
	code := m.Run()

	// Exit after tests
	os.Exit(code)
}

// TestCreateUser tests the creation of a new user (POST /users)
func TestCreateUser(t *testing.T) {
	// Create a user payload
	user := models.User{Name: "John Doe", Email: "johndoe@example.com", Age: 25}
	jsonUser, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Invoke the router
	router.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode the response body
	var createdUser models.User
	err = json.NewDecoder(rr.Body).Decode(&createdUser)
	if err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Check if the user was created correctly
	if createdUser.Name != user.Name || createdUser.Email != user.Email || createdUser.Age != user.Age {
		t.Errorf("Expected user data to match input, but got %+v", createdUser)
	}
}

// TestGetUser tests retrieving a user by ID (GET /users/{id})
func TestGetUser(t *testing.T) {
	// Create a request for user with ID 1 (assuming it exists)
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode the response body
	var user models.User
	err = json.NewDecoder(rr.Body).Decode(&user)
	if err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Check if user details are correct
	if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", user.ID)
	}
}

// TestUpdateUser tests updating a user by ID (PUT /users/{id})
func TestUpdateUser(t *testing.T) {
	// Create a payload for the update
	updatedUser := models.User{Name: "Jane Doe", Email: "janedoe@example.com", Age: 35}
	jsonUser, _ := json.Marshal(updatedUser)

	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if the status code is 204 No Content (successful update)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

// TestDeleteUser tests deleting a user by ID (DELETE /users/{id})
func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if the status code is 204 No Content (successful deletion)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

// TestGetNonExistingUser tests retrieving a non-existing user (GET /users/{id})
func TestGetNonExistingUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/9999", nil) // Assuming 9999 does not exist
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if status code is 404 Not Found
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// TestCreateUserWithInvalidData tests creating a user with invalid data (POST /users)
func TestCreateUserWithInvalidData(t *testing.T) {
	// Invalid user data (missing required fields like "name")
	user := models.User{Email: "invalid@example.com", Age: 25}
	jsonUser, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check if status code is 400 Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
