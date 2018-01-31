package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcome(t *testing.T) {
	// Create GET request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(welcome)

	// Run welcome function using created request and get result in rr - ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `<html>
<body>
   <br>
   <div>Welcome to Golang Predix Microservice Template</div>
   <br>
   <a href='/env' target='_blank'>Environment variables</a>
</body>
</html>`

	// Check response
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEnv(t *testing.T) {
	// Create GET request
	req, err := http.NewRequest("GET", "/env", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env)

	// Run env function using created request and get result in rr - ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `<div class="envs">`

	// Check response
	s := rr.Body.String()[:18]
	if s != expected {
	//if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}