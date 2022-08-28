package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAlice(t *testing.T) {
	result := HelloWorld("Alice")

	if result != "Hello Alice" {
		// error
		// t.Fail() // continue execute below code
		t.Error("Expected Hello Alice, but got", result)
	}

	fmt.Println("TestHelloWorldAlice Done")
}

func TestHelloWorldBob(t *testing.T) {
	result := HelloWorld("Bob")

	if result != "Hello Bob" {
		// error
		// t.FailNow() // stop immediately
		t.Fatal("Expected Hello Bob, but got", result)
	}

	fmt.Println("TestHelloWorldBob Done")
}
