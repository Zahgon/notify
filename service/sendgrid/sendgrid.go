package sendgrid

import (
	"context"

	"github.com/sendgrid/sendgrid-go"
)

// SendGrid struct holds necessary data to communicate with the SendGrid API.
type SendGrid struct {
	usePlainText      bool
	client            *sendgrid.Client
	senderAddress     string
	senderName        string
	receiverAddresses []string
}

// BodyType is used to specify the format of the body.
type BodyType int

const (
	// PlainText is used to specify that the body is plain text.
	PlainText BodyType = iota
	// HTML is used to specify that the body is HTML.
	HTML
)

// New returns a new instance of a SendGrid notification service.
// You will need a SendGrid API key.
// See https://sendgrid.com/docs/for-developers/sending-email/api-getting-started/
func New(apiKey, senderAddress, senderName string) *SendGrid { _ = "STUB: not implemented"; return nil }

// AddReceivers takes email addresses and adds them to the internal address list. The Send method will send
// a given message to all those addresses.
func (s *SendGrid) AddReceivers(addresses ...string) { _ = "STUB: not implemented"; return }

// BodyFormat can be used to specify the format of the body.
// Default BodyType is HTML.
func (s *SendGrid) BodyFormat(format BodyType) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set chats. Message body supports
// html as markup language.
func (s SendGrid) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new personalization instance to be able to add multiple receiver addresses.
