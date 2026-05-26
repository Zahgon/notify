package twitter

import (
	"context"

	"github.com/drswork/go-twitter/twitter"
)

// Twitter struct holds necessary data to communicate with the Twitter API.
type Twitter struct {
	client     *twitter.Client
	twitterIDs []string
}

// Credentials contains the authentication credentials needed for twitter
// api access
//
// ConsumerKey and ConsumerSecret can be thought of as the user name
// and password that represents your Twitter developer app when making
// API requests.
//
// An access token and access token secret are user-specific credentials
// used to authenticate OAuth 1.0a API requests.
// They specify the Twitter account the request is made on behalf of.
//
// See https://developer.twitter.com/en/docs/authentication/oauth-1-0a for more details.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// New returns a new instance of a Twitter service.
// For more information about Twitter access token:
//
//	-> https://developer.twitter.com/en/docs/authentication/oauth-1-0a/obtaining-user-access-tokens
func New(credentials Credentials) (*Twitter, error) { _ = "STUB: not implemented"; return nil, nil }

// Verify Credentials

// we can retrieve the user and verify if the credentials
// we have used successfully allow us to log in!

// AddReceivers takes TwitterIds and adds them to the internal twitterIDs list.
func (t *Twitter) AddReceivers(twitterIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set twitterIDs as a DM.
// See
// https://developer.twitter.com/en/docs/twitter-api/v1/direct-messages/sending-and-receiving/api-reference/new-event
func (t Twitter) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
