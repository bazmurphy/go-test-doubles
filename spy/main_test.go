package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// A "spy" is a variation of a stub that also records information about how it was called.
// ----------------------------------------------------------------------------

// The SpyMessageBroker struct has a slice publishedMessages to store the messages that were published to it.

// In the TestPublishMessage function, we create an instance of SpyMessageBroker called spyBroker.
// We then call PublishMessage twice with the spy broker and different messages.
// After publishing the messages, we use assert.Equal from the github.com/stretchr/testify/assert package to make assertions about the spy broker.
// We assert that the number of published messages is 2 and that the published messages match the expected messages in the correct order.

// The purpose of the spy message broker in this example is to record and verify the interactions with the message broker.
// By storing the published messages in the publishedMessages slice, we can track the messages that were sent to the broker.
// In the test, we can assert that the expected messages were published in the correct order.

// Spies are useful when you want to verify the behaviour of a system that interacts with external services or components.
// They allow you to record and inspect the interactions, ensuring that the system under test behaves as expected.

type SpyMessageBroker struct {
	publishedMessages []string
}

func (s *SpyMessageBroker) Publish(message string) {
	s.publishedMessages = append(s.publishedMessages, message)
}

func TestPublishMessage(t *testing.T) {
	t.Run("with spy message broker", func(t *testing.T) {
		message1 := "Hello, world!"
		message2 := "Goodbye, world!"

		spyBroker := &SpyMessageBroker{}

		PublishMessage(spyBroker, message1)
		PublishMessage(spyBroker, message2)

		assert.Equal(t, 2, len(spyBroker.publishedMessages), "unexpected number of published messages")
		assert.Equal(t, message1, spyBroker.publishedMessages[0], "unexpected published message")
		assert.Equal(t, message2, spyBroker.publishedMessages[1], "unexpected published message")
	})
}
