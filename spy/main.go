package main

import "fmt"

// ----------------------------------------------------------------------------
// A "spy" is a variation of a stub that also records information about how it was called.
// ----------------------------------------------------------------------------

// We define the MessageBroker interface and a struct RealMessageBroker.

// The RealMessageBroker simulates publishing a message to a real message broker.
// For simplicity, it just prints the message.

// The PublishMessage function takes a MessageBroker interface and a message string.
// It calls the Publish method of the provided message broker to publish the message.

// We demonstrate the usage of the real message broker.
// It creates an instance of RealMessageBroker and passes it to PublishMessage along with a message.

// In real-world usage, you would typically use the RealMessageBroker (or any other concrete message broker implementation) to publish messages to an actual message broker.
// The spy message broker is primarily used in testing to verify the behaviour of the code that depends on the message broker interface.

type MessageBroker interface {
	Publish(message string)
}

type RealMessageBroker struct{}

func (b *RealMessageBroker) Publish(message string) {
	// simulating publishing a message to a real message broker
	fmt.Printf("Publishing message: %s\n", message)
	// in a real implementation, this would publish the message to an actual message broker
}

func PublishMessage(broker MessageBroker, message string) {
	broker.Publish(message)
}

func main() {
	message := "Hello, world!"

	// using a real message broker
	realBroker := &RealMessageBroker{}
	PublishMessage(realBroker, message)
}
