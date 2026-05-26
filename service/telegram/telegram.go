package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	ModeMarkdown = tgbotapi.ModeMarkdown
	ModeHTML     = tgbotapi.ModeHTML
)

// HTML is the default mode.
//
//nolint:gochecknoglobals // I agree with the linter, won't bother fixing this now, will be fixed in v2.
var parseMode = ModeHTML

// Telegram struct holds necessary data to communicate with the Telegram API.
type Telegram struct {
	client  *tgbotapi.BotAPI
	chatIDs []int64
}

// New returns a new instance of a Telegram notification service.
// For more information about telegram api token:
//
//	-> https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api#NewBotAPI
func New(apiToken string) (*Telegram, error) { _ = "STUB: not implemented"; return nil, nil }

// SetClient set a new custom BotAPI instance.
// For example allowing you to use NewBotAPIWithClient:
//
//	-> https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api#NewBotAPIWithClient
func (t *Telegram) SetClient(client *tgbotapi.BotAPI) {
	_ = "STUB: not implemented"

	// SetParseMode sets the parse mode for the message body.
	// For more information about telegram constants:
	//
	//	-> https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api#pkg-constants
	return
}

func (t *Telegram) SetParseMode(mode string) {
	_ = "STUB: not implemented"

	// AddReceivers takes Telegram chat IDs and adds them to the internal chat ID list. The Send method will send
	// a given message to all those chats.
	return
}

func (t *Telegram) AddReceivers(chatIDs ...int64) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set chats. Message body supports
// html as markup language.
func (t Telegram) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
