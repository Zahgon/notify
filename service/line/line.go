package line

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Line struct holds info about client and destination ID for communicating with line API.
type Line struct {
	client      *linebot.Client
	receiverIDs []string
}

// New creates a new instance of Line notifier service
// For more info about line api credential:
// -> https://github.com/line/line-bot-sdk-go
func New(channelSecret, channelAccessToken string) (*Line, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers receives user, group or room IDs then add them to internal receivers list.
func (l *Line) AddReceivers(receiverIDs ...string) { _ = "STUB: not implemented"; return }

// Send receives message subject and body then sends it to all receivers set previously
// Subject will be on the first line followed by message on the next line.
func (l *Line) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
