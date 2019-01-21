package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Person Struct
// `json:...` is used to determine the names
// of the fields for the json
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Slice of persons
var people []Person

func main() {
	// Add a few Persons to the people slice
	people = append(people, Person{Name: "Cameron", Age: 18})
	people = append(people, Person{Name: "Jackson", Age: 19})
	people = append(people, Person{Name: "Ashley", Age: 33})
	people = append(people, Person{Name: "Clyde", Age: 41})

	// Set up the routes, and what functions they map to
	http.HandleFunc("/", home)
	http.HandleFunc("/people", api)

	// Listen and serve
	err := http.ListenAndServe(":8080", nil)

	// Check the error
	if err != nil {
		fmt.Printf("%v", err)
	}
}

// Function of type handler func
func home(w http.ResponseWriter, req *http.Request) {
	now := time.Now()

	// Utilizing the Fprint function which takes
	// a writer and a message
	fmt.Fprint(w, "<html><body>")
	// Fprintf takes a writer, message and
	// any extra paramters needed in the message
	fmt.Fprintf(w, "<h1>%s</h1>", now)
	fmt.Fprint(w, "<p>Go is kinda cool</p>")
	fmt.Fprint(w, "</body></html>")
}

func api(w http.ResponseWriter, req *http.Request) {
	// Check the method type
	if req.Method == "GET" {
		// Encode the people slice to JSON
		// And write it back out
		json.NewEncoder(w).Encode(people)

	} else if req.Method == "POST" {
		// Create an instance of a Person struct
		var new Person
		// Put the info from the request into the new person
		json.NewDecoder(req.Body).Decode(&new)
		// Add that person to the slice
		people = append(people, new)

		fmt.Fprint(w, "Success!")
	} else {
		fmt.Fprint(w, "Improper Request Method!")
	}
}
