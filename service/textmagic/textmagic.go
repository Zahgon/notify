package textmagic

import (
	"context"

	textMagic "github.com/textmagic/textmagic-rest-go-v2/v3"
)

// Service allow you to configure a TextMagic SDK client.
type Service struct {
	userName     string
	apiKey       string
	phoneNumbers []string
	client       *textMagic.APIClient
}

// New creates a new text magic client. Use your user-name and API key from
// https://my.textmagic.com/online/api/rest-api/keys.
func New(userName, apiKey string) *Service { _ = "STUB: not implemented"; return nil }

// AddReceivers adds the given phone numbers to the notifier.
func (s *Service) AddReceivers(phoneNumbers ...string) { _ = "STUB: not implemented"; return }

// Send sends a SMS via TextMagic to all previously added receivers.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
