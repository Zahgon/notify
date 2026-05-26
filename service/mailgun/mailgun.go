package mailgun

import (
	"context"

	"github.com/mailgun/mailgun-go/v5"
)

// Mailgun struct holds necessary data to communicate with the Mailgun API.
type Mailgun struct {
	client            *mailgun.Client
	domain            string
	senderAddress     string
	receiverAddresses []string
}

// New returns a new instance of a Mailgun notification service.
// You will need a Mailgun API key and domain name.
// See https://documentation.mailgun.com/en/latest/
func New(domain, apiKey, senderAddress string, opts ...Option) *Mailgun {
	_ = "STUB: not implemented"
	return nil
}

// AddReceivers takes email addresses and adds them to the internal address list. The Send method will send
// a given message to all those addresses.
func (m *Mailgun) AddReceivers(addresses ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set chats. Message body supports
// html as markup language.
func (m Mailgun) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
