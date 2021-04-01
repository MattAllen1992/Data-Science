// if you use the same package name in multiple files, they can
// see everything in each other's files, it essentially bridges them
package main

import (
	"fmt"      // formatted I/O
	"net/http" // implement HTTP clients and servers

	"github.com/gorilla/mux" // custom package to tailor HTTP routing (i.e. custom listen and send routes)
)

// main method (entry point into programme)
func main() {
	// create new mux router with routes
	r := newRouter()

	// tell the server to listen on port 8080 (this gives us an interface)
	// server now listens for all routes which have been registered to r
	http.ListenAndServe(":8080", r)
}

// constructor function for new routers
func newRouter() *mux.Router {
	// this line is the standard HTTP functionality, it is redundant here as we use MUX for custom routes instead
	// this method takes a path and a function
	// Go allows you to use functions as arguments to other functions
	// the only requirement is that the handler function has an appropriate signature (defined below)
	//http.HandleFunc("/", handler)

	// declare a new router
	r := mux.NewRouter()

	// declare specific methods that this custom path will be used for
	// registers a new route to the specified string/path and makes it a get request
	r.HandleFunc("/hello", handler).Methods("GET")

	// add API to get or create birds in the data
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	// point to static file directory
	staticFileDirectory := http.Dir("./assets/")

	// declare a handler which routes requests to the relevant filename
	// need to strip the prefix otherwise it reads it as "/assets/assets/"
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// create a matcher in our router which looks for all files in the "/assets/" filepath
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// return modified router
	return r
}

// this is our handler function, it takes an HTTP writer and request as inputs
// response writers generate HTTP messages whilst the request is a message itself
func handler(w http.ResponseWriter, r *http.Request) {
	// here we simply print "Hello World!" via our writer
	// Fprintf takes a writer object as its first argument and passes it the text/second argument
	fmt.Fprintf(w, "Hello World!")
}
