package mattermost

import (
	"context"

	"github.com/nikoksr/notify/service/http"
)

type httpClient interface {
	AddReceivers(wh ...*http.Webhook)
	PreSend(prefn http.PreSendHookFn)
	Send(ctx context.Context, subject, message string) error
	PostSend(postfn http.PostSendHookFn)
}

// Service encapsulates the notify httpService client and contains mattermost channel ids.
type Service struct {
	loginClient   httpClient
	messageClient httpClient
	channelIDs    map[string]bool
}

// New returns a new instance of a Mattermost notification service.
func New(url string) *Service { _ = "STUB: not implemented"; return nil }

// LoginWithCredentials provides helper for authentication using Mattermost user/admin credentials.
func (s *Service) LoginWithCredentials(ctx context.Context, loginID, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddReceivers takes Mattermost channel IDs or Chat IDs and adds them to the internal channel ID list.
// The Send method will send a given message to all these channels.
func (s *Service) AddReceivers(channelIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and send them to added channel ids.
// you will need a 'create_post' permission for your username.
// refer https://api.mattermost.com/ for more info.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// create post

// PreSend adds a pre-send hook to the service. The hook will be executed before sending a request to a receiver.
func (s *Service) PreSend(hook http.PreSendHookFn) { _ = "STUB: not implemented"; return }

// PostSend adds a post-send hook to the service. The hook will be executed after sending a request to a receiver.
func (s *Service) PostSend(hook http.PostSendHookFn) { _ = "STUB: not implemented"; return }

// setups main message service for creating posts.
func setupMsgService(url string) *http.Service {
	_ = "STUB: not implemented"
	// create new http client for sending messages/notifications
	return nil
}

// add custom payload builder

// add post-send hook for error checks

// setups login service to get token.
func setupLoginService(url string, msgService *http.Service) *http.Service {
	_ = "STUB: not implemented"
	// create another new http client for login request call.
	return nil
}

// append login path for the given mattermost server with custom payload builder.

// Add post-send hook to do error checks and log the response after it is received.
// Also extract token from response header and set it as part of pre-send hook of main http client for further
// requests.

// get token from header

// set token as pre-send hook
