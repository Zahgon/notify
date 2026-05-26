package lark

import (
	"context"
	"time"

	"github.com/go-lark/lark"

	"github.com/nikoksr/notify"
)

const defaultTimeout = 10 * time.Second

// CustomAppService is a Lark notify service using a Lark custom app.
type CustomAppService struct {
	receiveIDs []*ReceiverID
	cli        sendToer
}

// Compile time check that larkCustomAppService implements notify.Notifer.
var _ notify.Notifier = &CustomAppService{}

// NewCustomAppService returns a new instance of a Lark notify service using a
// Lark custom app.
func NewCustomAppService(appID, appSecret string) *CustomAppService {
	_ = "STUB: not implemented"
	return nil
}

// We need to set the bot to use Lark's open.larksuite.com domain instead of
// the default open.feishu.cn domain.

// Let the bot use a HTTP client with a longer timeout than the default 5
// seconds.

// AddReceivers adds recipients to future notifications. There are five different
// types of receiver IDs available in Lark and they must be specified here. For
// example:
//
//	larkService.AddReceivers(
//	  lark.OpenID("ou_c99c5f35d542efc7ee492afe11af19ef"),
//	  lark.UserID("8335aga2"),
//	  lark.UnionID("on_cad4860e7af114fb4ff6c5d496d1dd76"),
//	  lark.Email("xyz@example.com"),
//	  lark.ChatID("oc_a0553eda9014c201e6969b478895c230"),
//	)
func (c *CustomAppService) AddReceivers(ids ...*ReceiverID) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all
// previously registered recipient IDs.
func (c *CustomAppService) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// larkClientGoLarkChatBot is a wrapper around go-lark/lark's Bot, to be used
// for sending messages with custom apps.
type larkClientGoLarkChatBot struct {
	bot *lark.Bot
}

// SendTo implements the sendToer interface using a go-lark/lark chat bot.
func (l *larkClientGoLarkChatBot) SendTo(subject, message, receiverID, idType string) error {
	_ = "STUB: not implemented"
	return nil
}
