package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID         string // User ID (e.g., UUID)
	Username   string // User's username
	FirstName  string // User's first name
	LastName   string // User's last name
	Email      string // User's email address
	Password   string // User's password (make sure to hash and salt it in a real application)
	Age        int    // User's age
	IsAdmin    bool   // Indicates if the user is an admin
	// Add more fields as needed
}

// UserAPIRequest for GET request
type UserAPIGetRequest struct {
	ID string // User ID
}

// UserAPIRequest for POST request
type UserAPIPostRequest struct {
	Name       string // User's name
	Email      string // User's email
	ContactNo  string // User's contact number
	Password   string // User's password
	CollegeID  string // User's college ID (Foreign key)
	AccessLevel string // User's access level type
}

// UserAPIRequest for PUT request
type UserAPIPutRequest struct {
	ID         string // User ID
	Name       string // Updated name
	Email      string // Updated email
	ContactNo  string // Updated contact number
	Password   string // Updated password
	CollegeID  string // Updated college ID (Foreign key)
	AccessLevel string // Updated access level type
}

// UserAPIRequest for DELETE request
type UserAPIDeleteRequest struct {
	ID string // User ID to delete
}

// UserAPIResponse for all request methods
type UserAPIResponse struct {
	Message string // Response message
}

// Implement the Parse method for GET request
func (r *UserAPIGetRequest) Parse(req *http.Request) error {
	// Implement parsing logic for the GET request
	// Extract necessary data from the request and populate r
	r.ID = req.URL.Query().Get("id")
	return nil // Implement this as needed
}

// Implement the Parse method for POST request
func (r *UserAPIPostRequest) Parse(req *http.Request) error {
	// Implement parsing logic for the POST request
	// Extract necessary data from the request and populate r
	return json.NewDecoder(req.Body).Decode(r)
}

// Implement the Parse method for PUT request
func (r *UserAPIPutRequest) Parse(req *http.Request) error {
	// Implement parsing logic for the PUT request
	// Extract necessary data from the request and populate r
	r.ID = req.URL.Query().Get("id")
	return nil // Implement this as needed
}

// Implement the Parse method for DELETE request
func (r *UserAPIDeleteRequest) Parse(req *http.Request) error {
	// Implement parsing logic for the DELETE request
	// Extract necessary data from the request and populate r
	r.ID = req.URL.Query().Get("id")
	return nil // Implement this as needed
}

// Implement the Write method for UserAPIResponse
func (r *UserAPIResponse) Write(w http.ResponseWriter) error {
	// Implement serialization and writing logic for the User API response
	// Serialize the struct r and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}

// // Constructor for UserAPIGetRequest
// func NewUserAPIGetRequest(id string) *UserAPIGetRequest {
// 	return &UserAPIGetRequest{ID: id}
// }

// // Constructor for UserAPIPostRequest
// func NewUserAPIPostRequest(name, email, contactNo, password, collegeID, accessLevel string) *UserAPIPostRequest {
// 	return &UserAPIPostRequest{
// 		Name:       name,
// 		Email:      email,
// 		ContactNo:  contactNo,
// 		Password:   password,
// 		CollegeID:  collegeID,
// 		AccessLevel: accessLevel,
// 	}
// }

// // Constructor for UserAPIPutRequest
// func NewUserAPIPutRequest(id, name, email, contactNo, password, collegeID, accessLevel string) *UserAPIPutRequest {
// 	return &UserAPIPutRequest{}
// }

// // Constructor for UserAPIDeleteRequest
// func NewUserAPIDeleteRequest(id string) *UserAPIDeleteRequest {
// 	return &UserAPIDeleteRequest{}
// }
