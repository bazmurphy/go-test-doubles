package main

import (
	"testing"
)

// TestProcessData tests the ProcessData function with both the real logger and the dummy logger.

// The purpose of the dummy logger in this example is to satisfy the requirement of providing a logger object to the ProcessData function, even though the logger is not actually used within the function.
// It allows you to test the behaviour of ProcessData without relying on the logging functionality.

func TestProcessData(t *testing.T) {
	data := "test data"

	// "with real logger":
	// It creates an instance of RealLogger and passes it to ProcessData.
	// No assertions are needed here because the real logger is used, and its behaviour is not being tested.
	t.Run("with real logger", func(t *testing.T) {
		realLogger := &RealLogger{}
		ProcessData(data, realLogger)
	})

	// "with dummy logger":
	// It creates an instance of DummyLogger and passes it to ProcessData.
	// Again, no assertions are needed because the dummy logger doesn't do anything.
	t.Run("with dummy logger", func(t *testing.T) {
		dummyLogger := &DummyLogger{}
		ProcessData(data, dummyLogger)
	})
}
