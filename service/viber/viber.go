package viber

import (
	"context"

	vb "github.com/mileusna/viber"
)

type viberClient interface {
	SetWebhook(url string, eventTypes []string) (vb.WebhookResp, error)
	SendTextMessage(receiver, msg string) (uint64, error)
}

// Compile-time check to ensure that vb.Viber implements the viberClient interface.
var _ viberClient = new(vb.Viber)

// Viber struct holds necessary fields to communicate with Viber API.
type Viber struct {
	Client            viberClient
	SubscribedUserIDs []string
}

// New returns a new instance of Viber notification service.
func New(appKey, senderName, senderAvatar string) *Viber { _ = "STUB: not implemented"; return nil }

// AddReceivers receives subscribed user IDs then add them to internal receivers list.
func (v *Viber) AddReceivers(subscribedUserIDs ...string) { _ = "STUB: not implemented"; return }

// SetWebhook receives a URL that will we used as a webhook URL for Viber.
func (v *Viber) SetWebhook(webhookURL string) error { _ = "STUB: not implemented"; return nil }

// Send takes a message subject and a message body and sends them to all previously set userIds.
func (v *Viber) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
