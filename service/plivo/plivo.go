package plivo

import (
	"context"
	"net/http"

	plivo "github.com/plivo/plivo-go/v7"
)

// ClientOptions allow you to configure a Plivo SDK client.
type ClientOptions struct {
	AuthID    string // If empty, env variable PLIVO_AUTH_ID will be used
	AuthToken string // If empty, env variable PLIVO_AUTH_TOKEN will be used

	// Optional
	HTTPClient *http.Client // Bring Your Own Client
}

// MessageOptions allow you to configure options for sending a message.
type MessageOptions struct {
	Source string // a Plivo source phone number or a Plivo Powerpack UUID

	// Optional
	CallbackURL    string // URL to which status update callbacks for the message should be sent
	CallbackMethod string // The HTTP method to be used when calling CallbackURL - GET or POST(default)
}

// plivoMsgClient abstracts Plivo SDK for writing unit tests.
type plivoMsgClient interface {
	Create(plivo.MessageCreateParams) (*plivo.MessageCreateResponseBody, error)
}

// Service is a Plivo client.
type Service struct {
	client       plivoMsgClient
	mopts        MessageOptions
	destinations []string
}

// New creates a new instance of plivo service.
func New(cOpts *ClientOptions, mOpts *MessageOptions) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers adds the given destination phone numbers to the notifier.
func (s *Service) AddReceivers(phoneNumbers ...string) { _ = "STUB: not implemented"; return }

// Send sends a SMS via Plivo to all previously added receivers.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// multiple destinations, use bulk message syntax
// see: https://www.plivo.com/docs/sms/api/message#bulk-messaging
