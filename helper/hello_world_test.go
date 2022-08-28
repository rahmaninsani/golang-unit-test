package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Skip test on darwin/macOS")
	}

	result := HelloWorld("Charlie")
	errorMessage := fmt.Sprintf("Expected 'Hello Charlie', but got '%s'", result)
	require.Equal(t, "Hello Charlie", result, errorMessage)
	fmt.Println("TestHelloWorldRequire Done")
}

func TestMain(m *testing.M) {
	// Before -> e.g: connect to database
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Alice", func(t *testing.T) {
		result := HelloWorld("Alice")
		errorMessage := fmt.Sprintf("Expected 'Hello Alice', but got '%s'", result)
		require.Equal(t, "Hello Alice", result, errorMessage)
	})

	t.Run("Bob", func(t *testing.T) {
		result := HelloWorld("Bob")
		errorMessage := fmt.Sprintf("Expected 'Hello Bob', but got '%s'", result)
		require.Equal(t, "Hello Bob", result, errorMessage)
	})
}
