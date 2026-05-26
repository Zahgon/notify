package msteams

import (
	"context"
	"net/http"

	teams "github.com/atc0005/go-teams-notify/v2"
)

type teamsClient interface {
	// https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2#TeamsClient.SendWithContext
	SendWithContext(ctx context.Context, webhookURL string, message teams.TeamsMessage) error
	// https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2#TeamsClient.SkipWebhookURLValidationOnSend
	SkipWebhookURLValidationOnSend(skip bool) *teams.TeamsClient
	// https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2#TeamsClient.SetHTTPClient
	SetHTTPClient(httpClient *http.Client) *teams.TeamsClient
	// https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2#TeamsClient.SetUserAgent
	SetUserAgent(userAgent string) *teams.TeamsClient
}

// Compile-time check to ensure that teams.Client implements the teamsClient interface.
var _ teamsClient = teams.NewTeamsClient()

// MSTeams struct holds necessary data to communicate with the MSTeams API.
type MSTeams struct {
	client   teamsClient
	webHooks []string

	wrapText bool
}

// New returns a new instance of a MSTeams notification service.
// For more information about telegram api token:
//
//	-> https://github.com/atc0005/go-teams-notify#example-basic
func New() *MSTeams { _ = "STUB: not implemented"; return nil }

// DisableWebhookValidation disables the validation of webhook URLs, including the validation of known prefixes so that
// custom/private webhook URL endpoints can be used (e.g., testing purposes).
// For more information about telegram api token:
//
//	-> https://github.com/atc0005/go-teams-notify#example-disable-webhook-url-prefix-validation
func (m *MSTeams) DisableWebhookValidation() { _ = "STUB: not implemented"; return }

// WithWrapText sets the wrapText field to the provided value. This is disabled by default.
func (m *MSTeams) WithWrapText(wrapText bool) { _ = "STUB: not implemented"; return }

// AddReceivers takes MSTeams channel web-hooks and adds them to the internal web-hook list. The Send method will send
// a given message to all those chats.
func (m *MSTeams) AddReceivers(webHooks ...string) { _ = "STUB: not implemented"; return }

// SetUseragent allows the user to set a custom user agent.
func (m *MSTeams) SetUseragent(userAgent string) { _ = "STUB: not implemented"; return }

// SetHTTPClient allows the user to set a custom http client.
func (m *MSTeams) SetHTTPClient(httpClient *http.Client) { _ = "STUB: not implemented"; return }

// Send accepts a subject and a message body and sends them to all previously specified channels. Message body supports
// html as markup language.
// For more information about telegram api token:
//
//	-> https://github.com/atc0005/go-teams-notify#example-basic
func (m MSTeams) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
