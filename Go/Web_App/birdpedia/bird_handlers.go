package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// create Bird struct with species and description attributes
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

// create array of Bird objects
// we'll use this throughout for handling birds
var birds []Bird

// handler function to get existing bird data
func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// convert variable to json
	birdListBytes, err := json.Marshal(birds)

	// if error, send message to user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// otherwise return json list of birds as response
	w.Write(birdListBytes)
}

// handler function to create new bird data
func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// create empty instance of Bird
	bird := Bird{}

	// parse form values as HTML
	// this populates r.Form (see below) with the contents of the HTML
	err := r.ParseForm()

	// check for errors and handle accordingly
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// extract bird info from parsed form
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// append data to existing birds list
	birds = append(birds, bird)

	// re-direct user to HTML assets page
	// this receives the request (including URL) from the writer and request
	// it then redirects them using the string "/assets/" (with the original request URL)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
