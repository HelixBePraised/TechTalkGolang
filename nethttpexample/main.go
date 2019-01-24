package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Utilizing the Fprint function which takes
	// a writer and a message
	_, err := fmt.Fprint(w, `<html>
						<body>
							<h1>API Usage</h1>
							<p><bold>/people - GET</bold> Used to list out people.</p>
							<p>
								<bold>/people - POST</bold> Used to add a person to list of people	
								Fields Required:
								<ul>
									<li>name - string - Name of person</li>
									<li>age - int - Age of person</li>
								</ul>
							</p>
						</body>
						</html>
					`)
	if err != nil {
		fmt.Printf("Error writing to response: %v", err)
	}
}

func api(w http.ResponseWriter, req *http.Request) {
	// Check the method type
	if req.Method == "GET" {
		// Encode the people slice to JSON
		// And write it back out
		encoder := json.NewEncoder(w)
		err := encoder.Encode(people)

		if err != nil {
			fmt.Printf("Something went wrong encoding data: %v", err)
		}

	} else if req.Method == "POST" {
		// Create an instance of a Person struct
		var newPerson Person
		// Put the info from the request into the new person
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&newPerson)
		if err != nil {
			fmt.Printf("Error decoding data: %v", err)
		}
		// Add that person to the slice
		people = append(people, newPerson)

		_, err = fmt.Fprint(w, "Success!")
		if err != nil {
			fmt.Printf("Error writing back to response: %v", err)
		}
	} else {
		_, err := fmt.Fprint(w, "Improper Request Method!")
		if err != nil {
			fmt.Printf("Error writin back to response: %v", err)
		}
	}
}
