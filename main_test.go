package main

import (
	"testing"
)

func TestMainPr(t *testing.T) {

	// Check the response body is what we expect.
	expected := "I do stuff!"
	if doStuff() != expected {
		t.Errorf("doStuff(): got %v want %v",
			doStuff(), expected)
	}

}
