package slack

import (
	"context"

	"github.com/slack-go/slack"
)

type slackClient interface {
	PostMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, error)
}

// Compile-time check to ensure that slack.Client implements the slackClient interface.
var _ slackClient = new(slack.Client)

// Slack struct holds necessary data to communicate with the Slack API.
type Slack struct {
	client     slackClient
	channelIDs []string
}

// New returns a new instance of a Slack notification service.
// For more information about slack api token:
//
//	-> https://pkg.go.dev/github.com/slack-go/slack#New
func New(apiToken string) *Slack { _ = "STUB: not implemented"; return nil }

// AddReceivers takes Slack channel IDs and adds them to the internal channel ID list. The Send method will send
// a given message to all those channels.
func (s *Slack) AddReceivers(channelIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set channels.
// you will need a slack app with the chat:write.public and chat:write permissions.
// see https://api.slack.com/
func (s Slack) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
