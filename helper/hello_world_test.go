package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bob")
	if result != "Hello Bob" {
		// error
		panic("Expected 'Hello Bob', got " + result)
	}
}
