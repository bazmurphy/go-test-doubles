package main

import (
	"fmt"
	"log"
)

// A "dummy" is an object that is passed around but never actually used.

// We define the Logger interface and two structs: RealLogger and DummyLogger.

// The RealLogger provides a concrete implementation of the Log method, which logs the message using log.Println.

// The DummyLogger has an empty implementation of the Log method, which does nothing.

type Logger interface {
	Log(message string)
}

type RealLogger struct{}

func (rl *RealLogger) Log(message string) {
	log.Println("Real Logger:", message)
}

// this is "dummy" an object that is passed around but never used
type DummyLogger struct{}

func (dl *DummyLogger) Log(message string) {}

// The ProcessData function takes a data string and a logger of type Logger.
// It calls the Log method of the provided logger to log a message and then performs the data processing logic.

func ProcessData(data string, logger Logger) {
	logger.Log("Processing data: " + data)
	// perform data processing logic here
	fmt.Println("Data processed:", data)
}

// In this example, we simply print a message indicating that the data has been processed.
// When using the RealLogger, the logged message is printed to the console.

func main() {
	data := "some data"

	// using the real logger
	realLogger := &RealLogger{}
	ProcessData(data, realLogger)

	// using the dummy logger (! typically only used in the tests)
	dummyLogger := &DummyLogger{}
	ProcessData(data, dummyLogger)

	// In real-world usage, you would typically use the RealLogger (or any other concrete logger implementation) to perform actual logging.
	// The dummy logger is primarily used in testing or in situations where logging is not required but the function expects a logger object.
}
