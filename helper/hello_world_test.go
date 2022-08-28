package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// assert -> Fail()
// require -> FailNow()

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Charlie")
	errorMessage := fmt.Sprintf("Expected 'Hello Charlie', but got '%s'", result)
	assert.Equal(t, "Hello Charlie", result, errorMessage)
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Charlie")
	errorMessage := fmt.Sprintf("Expected 'Hello Charlie', but got '%s'", result)
	require.Equal(t, "Hello Charlie", result, errorMessage)
	fmt.Println("TestHelloWorldRequire Done")
}

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
