package main

// ----------------------------------------------------------------------------
// A "mock" object acts as a stand-in for a real dependency, and it helps you define and verify the expected interactions between the code you're testing and that dependency
// ----------------------------------------------------------------------------

// We define the EmailSender interface and a struct RealEmailSender.

// The RealEmailSender simulates sending an email using a real email service. For simplicity, it just prints the recipient, subject, and body of the email.

// The SendWelcomeEmail function takes an EmailSender interface and a recipient string. It composes the subject and body of the welcome email and calls the SendEmail method of the provided email sender.

// We demonstrate the usage of the real email sender.
// It creates an instance of RealEmailSender and passes it to SendWelcomeEmail along with the recipient email address.

// In real-world usage, you would typically use the RealEmailSender (or any other concrete email sender implementation) to send actual emails.
// The mock email sender is primarily used in testing to verify the behaviour of the code that depends on the email sender interface.

import (
	"fmt"
)

type EmailSender interface {
	SendEmail(recipient, subject, body string) error
}

type RealEmailSender struct{}

func (s *RealEmailSender) SendEmail(recipient, subject, body string) error {
	// simulating sending an email using a real email service
	fmt.Printf("Sending email | recipient: %s | subject: %s | body: %s\n", recipient, subject, body)
	// in a real implementation, this would send the email using an actual email service
	return nil
}

func SendWelcomeEmail(emailSender EmailSender, recipient string) error {
	subject := "Welcome to our app"
	body := "Thank you for signing up!"
	return emailSender.SendEmail(recipient, subject, body)
}

func main() {
	recipient := "real_email@example.com"

	// using a real email sender
	realEmailSender := &RealEmailSender{}
	err := SendWelcomeEmail(realEmailSender, recipient)
	if err != nil {
		fmt.Printf("error sending welcome email: %v\n", err)
	}
}
