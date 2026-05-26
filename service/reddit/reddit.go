package reddit

import (
	"context"

	"github.com/caarlos0/go-reddit/v3/reddit"
)

type redditMessageClient interface {
	Send(context.Context, *reddit.SendMessageRequest) (*reddit.Response, error)
}

// Compile-time check to ensure that reddit.MessageService implements the redditMessageClient interface.
var _ redditMessageClient = new(reddit.MessageService)

// Reddit struct holds necessary data to communicate with the Reddit API.
type Reddit struct {
	client     redditMessageClient
	recipients []string
}

// New returns a new instance of a Reddit notification service.
// For more information on obtaining client credentials:
//
//	-> https://github.com/reddit-archive/reddit/wiki/OAuth2
func New(clientID, clientSecret, username, password string) (*Reddit, error) {
	_ = "STUB: not implemented"
	// Disable HTTP2 in http client
	// Details:
	// https://www.reddit.com/r/redditdev/comments/t8e8hc/getting_nothing_but_429_responses_when_using_go/i18yga2/
	return nil, nil
}

// AddReceivers takes Reddit usernames and adds them to the internal recipient list. The Send method will send
// a given message to all of those users.
func (r *Reddit) AddReceivers(recipients ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set recipients.
func (r *Reddit) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
