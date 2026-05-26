package twilio

import (
	"context"
	"net/url"

	"github.com/kevinburke/twilio-go"
)

// Compile-time check that twilio.MessageService satisfies twilioClient interface.
var _ twilioClient = &twilio.MessageService{}

// twilioClient abstracts twilio-go MessageService for writing unit tests.
type twilioClient interface {
	SendMessage(from, to, body string, mediaURLs []*url.URL) (*twilio.Message, error)
}

// Service encapsulates the Twilio Message Service client along with internal state for storing recipient phone numbers.
type Service struct {
	client twilioClient

	fromPhoneNumber string
	toPhoneNumbers  []string
}

// New returns a new instance of Twilio notification service.
func New(accountSID, authToken, fromPhoneNumber string) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes strings of recipient phone numbers and appends them to the internal phone numbers slice.
// The Send method will send a given message to all those phone numbers.
func (s *Service) AddReceivers(phoneNumbers ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set phone numbers.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
