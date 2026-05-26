package pushover

import (
	"context"

	"github.com/gregdel/pushover"
)

type pushoverClient interface {
	SendMessage(*pushover.Message, *pushover.Recipient) (*pushover.Response, error)
}

// Compile-time check to ensure that pushover.Pushover implements the pushoverClient interface.
var _ pushoverClient = new(pushover.Pushover)

// Pushover struct holds necessary data to communicate with the Pushover API.
type Pushover struct {
	client     pushoverClient
	recipients []pushover.Recipient
}

// New returns a new instance of a Pushover notification service.
// For more information about Pushover app token:
//
//	-> https://support.pushover.net/i175-how-do-i-get-an-api-or-application-token
func New(appToken string) *Pushover { _ = "STUB: not implemented"; return nil }

// AddReceivers takes Pushover user/group IDs and adds them to the internal recipient list. The Send method will send
// a given message to all of those recipients.
func (p *Pushover) AddReceivers(recipientIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set recipients.
func (p Pushover) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
