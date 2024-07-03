package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// ----------------------------------------------------------------------------
// A "mock" object acts as a stand-in for a real dependency, and it helps you define and verify the expected interactions between the code you're testing and that dependency
// ----------------------------------------------------------------------------

// The MockEmailSender struct embeds the mock.Mock type from the github.com/stretchr/testify/mock package.
// It defines the SendEmail method that records the arguments passed to it and returns the mocked error value.

type MockEmailSender struct {
	mock.Mock
}

func (m *MockEmailSender) SendEmail(recipient, subject, body string) error {
	fmt.Printf("sending mock email | recipient: %s | subject: %s | body: %s\n", recipient, subject, body)
	mockArguments := m.Called(recipient, subject, body)
	// fmt.Printf("DEBUG | MockEmailSender | SendEmail | mockArguments: %v\n", mockArguments)
	// fmt.Printf("DEBUG | MockEmailSender | SendEmail | mockArguments.Error(0): %v\n", mockArguments.Error(0))
	return mockArguments.Error(0)
}

// In the TestSendWelcomeEmail function, we create an instance of MockEmailSender and set up expectations using mockEmailSender.On("SendEmail", recipient, subject, body).Return(nil).
// This specifies that when the SendEmail method is called with the provided arguments, it should return nil, indicating no error.

// We then call SendWelcomeEmail with the mock email sender and the recipient email address.

// After the function call, we use mockEmailSender.AssertExpectations(t) to verify that all the expected method calls were made on the mock object.
// If any expected call is missing or has different arguments, the test will fail.

// The purpose of the mock email sender in this example is to verify the behaviour of the SendWelcomeEmail function and ensure that it interacts with the email sender correctly.
// By setting up expectations on the mock object, we can specify the expected method calls and their return values.
// This allows us to test the function in isolation, without actually sending real emails.

func TestSendWelcomeEmail(t *testing.T) {
	t.Run("with mock email sender", func(t *testing.T) {
		recipient := "test_email@test.com"
		subject := "Welcome to our app"
		body := "Thank you for signing up!"

		mockEmailSender := new(MockEmailSender)
		mockEmailSender.On("SendEmail", recipient, subject, body).Return(nil)

		err := SendWelcomeEmail(mockEmailSender, recipient)
		mockEmailSender.AssertExpectations(t)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
