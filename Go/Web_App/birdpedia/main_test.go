package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Go let's you create files called *_test.go and it will test files with matching
// names (before the underscore) using the inputs we'll provide here
func TestHandler(t *testing.T) {
	// here we create a custom request which we will pass to our handler to test it
	// that args here are method (i.e. message type), route string we're searching for
	// and request body (only if looking for specific contents)
	// i.e. get request at specific path with specific content
	req, err := http.NewRequest("GET", "", nil)

	// handle unexpected errors
	if err != nil {
		t.Fatal(err)
	}

	// httptest allows us to create a dummy browser that we can send our request to
	// this helps us simulate a real world HTTP send/receive process for testing
	recorder := httptest.NewRecorder()

	// create an HTTP handler using the function we defined in main.go
	hf := http.HandlerFunc(handler)

	// serve HTTP request to our recorder (executes our handler)
	hf.ServeHTTP(recorder, req)

	// compare expected to actual and flag errors
	expected := "Hello World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// test router for expected route
func TestRouter(t *testing.T) {
	// create custom router using constructor in main.go
	r := newRouter()

	// create a new server (auto-closes once messages stop)
	mockServer := httptest.NewServer(r)

	// make a get request to our server to the custom "/hello" route defined in our router
	resp, err := http.Get(mockServer.URL + "/hello")

	// handle unexpected errors
	if err != nil {
		t.Fatal(err)
	}

	// we expect a status code of OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be OK:200, received %d", resp.StatusCode)
	}

	// close response body after following code completes
	defer resp.Body.Close()

	// read response body into var (check for errors)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// convert to string
	respString := string(b)
	expected := "Hello World!"

	// check if actual response matches expected
	if respString != expected {
		t.Errorf("Response should be %v, instead received %v", expected, respString)
	}
}

// test non-existent route (should fail with expected status code)
func TestRouterForNonExistentRoute(t *testing.T) {
	// create new router (with "/hello" route)
	r := newRouter()

	// create a server with r acting as the middle-man between this app and the server
	mockServer := httptest.NewServer(r)

	// make a get request which doesn't match "/hello"
	resp, err := http.Get(mockServer.URL + "/goodbye")

	// handle unexpected errors
	if err != nil {
		t.Fatal(err)
	}

	// this should fail with code 405
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, received %d", resp.StatusCode)
	}

	// we will also check the response body, we expect it to be empty
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respBody := string(b)
	expected := ""
	if respBody != expected {
		t.Errorf("Expected empty response \"%v\", received %v", expected, respBody)
	}
}

// test static file handling
func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// get request sent to static assets folder
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// should get 200 status for OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, instead received %v", resp.StatusCode)
	}

	// testing the entire content of an html file could be dangerous if it's huge
	// therefore we'll just test the header to check it is an html file that's been served
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Expected content type %v, instead received %v", expectedContentType, contentType)
	}
}

// test get bird data
func TestGetBirdHandler(t *testing.T) {

}
