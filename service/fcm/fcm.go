package fcm

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/messaging"

	"github.com/nikoksr/notify"
)

// Compile-time check that Service satisfies the Notifier interface.
var _ notify.Notifier = (*Service)(nil)

// Used to generate mocks for the FCM client.
type fcmClient interface {
	Send(ctx context.Context, message ...*messaging.Message) (*messaging.BatchResponse, error)
	SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
}

// Service encapsulates the FCM client along with internal state for storing device tokens.
type Service struct {
	client       fcmClient
	deviceTokens []string
}

// Option is a function that configures a Service.
type Option func(*Service) error

// WithCredentialsFile returns an Option to configure the FCM client with a credentials file.
func WithCredentialsFile(filename string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithProjectID returns an Option to configure the FCM client with a project ID.
func WithProjectID(projectID string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient returns an Option to configure the FCM client with a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// New returns a new instance of a FCM notification service.
func New(ctx context.Context, opts ...Option) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes FCM device tokens and appends them to the internal device tokens slice.
func (s *Service) AddReceivers(deviceTokens ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set devices.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
