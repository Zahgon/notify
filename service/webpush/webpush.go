//nolint:gochecknoglobals // I agree with the linter, won't bother fixing this now, will be fixed in v2.
package webpush

import (
	"context"

	"github.com/SherClockHolmes/webpush-go"
)

type (
	// Urgency indicates the importance of the message. It's a type alias for webpush.Urgency.
	Urgency = webpush.Urgency

	// Options are optional settings for the sending of a message. It's a type alias for webpush.Options.
	Options = webpush.Options

	// Subscription is a JSON representation of a webpush subscription. It's a type alias for webpush.Subscription.
	Subscription = webpush.Subscription

	// messagePayload is the JSON payload that is sent to the webpush endpoint.
	messagePayload struct {
		Subject string         `json:"subject"`
		Message string         `json:"message"`
		Data    map[string]any `json:"data,omitempty"`
	}

	msgDataKey    struct{}
	msgOptionsKey struct{}
)

// optionsKey is used as a context.Context key to optionally add options to the messagePayload payload.
var optionsKey = msgOptionsKey{}

// dataKey is used as a context.Context key to optionally add data to the messagePayload payload.
var dataKey = msgDataKey{}

// These are exposed Urgency constants from the webpush package.
var (
	// UrgencyVeryLow requires device state: on power and Wi-Fi.
	UrgencyVeryLow Urgency = webpush.UrgencyVeryLow

	// UrgencyLow requires device state: on either power or Wi-Fi.
	UrgencyLow Urgency = webpush.UrgencyLow

	// UrgencyNormal excludes device state: low battery.
	UrgencyNormal Urgency = webpush.UrgencyNormal

	// UrgencyHigh admits device state: low battery.
	UrgencyHigh Urgency = webpush.UrgencyHigh
)

// Service encapsulates the webpush notification system along with the internal state.
type Service struct {
	subscriptions []webpush.Subscription
	options       webpush.Options
}

// New returns a new instance of the Service.
func New(vapidPublicKey string, vapidPrivateKey string) *Service {
	_ = "STUB: not implemented"
	return nil
}

// AddReceivers adds one or more subscriptions to the Service.
func (s *Service) AddReceivers(subscriptions ...Subscription) { _ = "STUB: not implemented"; return }

// withOptions returns a new Options struct with the incoming options merged with the Service's options. The incoming
// options take precedence, except for the VAPID keys. Existing VAPID keys are only replaced if the incoming VAPID keys
// are not empty.
func (s *Service) withOptions(options Options) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

// WithOptions binds the options to the context so that they will be used by the Service.Send method automatically.
// Options
// are settings that allow you to customize the sending behavior of a message.
func WithOptions(ctx context.Context, options Options) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func optionsFromContext(ctx context.Context) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

// WithData binds the data to the context so that it will be used by the Service.Send method automatically. Data is a
// map[string]any and acts as a metadata field that is sent along with the message payload.
func WithData(ctx context.Context, data map[string]any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func dataFromContext(ctx context.Context) map[string]any { _ = "STUB: not implemented"; return nil }

// payloadFromContext returns a json encoded byte array of the messagePayload payload that is ready to be sent to the
// webpush endpoint. Internally, it uses the messagePayload and data from the context, and it combines it with the
// subject and message arguments into a single messagePayload.
func payloadFromContext(ctx context.Context, subject, message string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load optional data

// send is a wrapper that makes it primarily easier to defer the closing of the response body.
func (s *Service) send(ctx context.Context, message []byte, subscription *Subscription, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

// Everything is fine

// Make sure to produce a helpful error message

// Send sends a message to all the webpush subscriptions that have been added to the Service. The subject and message
// arguments are the subject and message of the messagePayload payload. The context can be used to optionally add
// options and data to the messagePayload payload. See the WithOptions and WithData functions.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	// Get the options from the context and merge them with the service's initial options
	return nil
}
