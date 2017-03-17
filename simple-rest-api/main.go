package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person - A person.
type Person struct {
	ID        int      `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address - A persons address.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// All the people, so many people.
var people []Person

// Port number.
const port = 12345

// GetPeopleEndpoint - Get all people.
//
// Params:
//     w http.ResponseWriter - The HTTP response.
//     req *http.Request - The HTTP request.
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	// Set the content type header.
	w.Header().Set("Content-Type", "application/json")

	// Set the response.
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint - Get a person.
//
// Params:
//     w http.ResponseWriter - The HTTP response.
//     req *http.Request - The HTTP request.
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get the URL parameters.
	params := mux.Vars(req)

	// Get the ID parameter as an integer.
	reqID, _ := strconv.Atoi(params["id"])

	// Check for a matching person.
	for _, person := range people {
		if person.ID == reqID {
			// Set the content type header.
			w.Header().Set("Content-Type", "application/json")

			// Set the response.
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	// Throw a 404 if we cannot find that person.
	http.NotFound(w, req)
}

// CreatePersonEndpoint - Create a person.
//
// Params:
//     w http.ResponseWriter - The HTTP response.
//     req *http.Request - The HTTP request.
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Create a new person.
	var person Person

	// Populate the new person from the JSON body.
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		// Set an error code.
		w.WriteHeader(http.StatusBadRequest)

		// Set an error messsage.
		errorMessage := fmt.Sprintf("Could not create new person. Reason: %+v", err)
		w.Write([]byte(errorMessage))
		return
	}

	// Set the person ID.
	person.ID = GetNextPersonID()

	// Add to the list of people.
	people = append(people, person)

	// Set the content type header.
	w.Header().Set("Content-Type", "application/json")

	// Set the response.
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint - Delete a person.
//
// Params:
//     w http.ResponseWriter - The HTTP response.
//     req *http.Request - The HTTP request.
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get the URL parameters.
	params := mux.Vars(req)

	// Get the ID parameter as an integer.
	reqID, _ := strconv.Atoi(params["id"])

	// Check for a matching person.
	for index, person := range people {
		if person.ID == reqID {
			// Delete the person.
			people = append(people[:index], people[index+1:]...)

			// Set an status messsage.
			statusMessage := fmt.Sprintf("Person (%d) was deleted", reqID)
			w.Write([]byte(statusMessage))
			return
		}
	}

	// Throw a 404 if we cannot find that person.
	http.NotFound(w, req)
}

// GetNextPersonID - Get the next available person ID.
//
// Return:
//     int - The next available person ID.
func GetNextPersonID() int {
	var maxPersonID int
	for _, person := range people {
		if person.ID > maxPersonID {
			maxPersonID = person.ID
		}
	}

	return maxPersonID + 1
}

func main() {
	// Add some example people. In reality this would come from a database.
	people = append(people, Person{ID: 1, Firstname: "Dan", Lastname: "Richards", Address: &Address{City: "Bournemouth", State: "Dorset"}})
	people = append(people, Person{ID: 2, Firstname: "Emma", Lastname: "Richards"})

	// Create the router.
	router := mux.NewRouter()

	// Routes.
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	// Serve!
	fmt.Printf("Serving on port %d. Press CTRL+C to cancel.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
