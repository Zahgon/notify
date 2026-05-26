package lark

import (
	"context"

	"github.com/go-lark/lark"

	"github.com/nikoksr/notify"
)

// WebhookService is a Notify service that uses a Lark webhook to send messages.
type WebhookService struct {
	cli sender
}

// Compile time check that larkCustomAppService implements notify.Notifer.
var _ notify.Notifier = &WebhookService{}

// NewWebhookService returns a new instance of a Lark notify service using a
// Lark group chat webhook. Note that this service does not take any
// notification receivers because it can only push messages to the group chat
// it belongs to.
func NewWebhookService(webhookURL string) *WebhookService { _ = "STUB: not implemented"; return nil }

// Send sends the message subject and body to the group chat.
func (w *WebhookService) Send(_ context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// larkClientGoLarkNotificationBot is a wrapper around go-lark/lark's Bot, to
// be used for notifications via webhooks only.
type larkClientGoLarkNotificationBot struct {
	bot *lark.Bot
}

// Send implements the sender interface using a go-lark/lark notification bot.
func (w *larkClientGoLarkNotificationBot) Send(subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
