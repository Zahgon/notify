package bark

import (
	"context"
	"net/http"
)

// Service allow you to configure Bark service.
type Service struct {
	deviceKey  string
	client     *http.Client
	serverURLs []string
}

func defaultHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

//nolint: mnd // 5 seconds is a reasonable timeout for a push notification

// DefaultServerURL is the default server to use for the bark service.
const DefaultServerURL = "https://api.day.app/"

// normalizeServerURL normalizes the server URL. It prefixes it with https:// if it's not already and appends a slash
// if it's not already there. If the serverURL is empty, the DefaultServerURL is used. We're not validating the url here
// on purpose, we leave that to the http client.
func normalizeServerURL(serverURL string) string { _ = "STUB: not implemented"; return "" }

// Normalize the url

// AddReceivers adds server URLs to the list of servers to use for sending messages. We call it Receivers and not
// servers because strictly speaking, the server is still receiving the message, and additionally we're following the
// naming convention of the other services.
func (s *Service) AddReceivers(serverURLs ...string) { _ = "STUB: not implemented"; return }

// NewWithServers returns a new instance of Bark service. You can use this service to send messages to bark. You can
// specify the servers to send the messages to. By default, the service will use the default server
// (https://api.day.app/) if you don't specify any servers.
func NewWithServers(deviceKey string, serverURLs ...string) *Service {
	_ = "STUB: not implemented"
	return nil
}

// Calling service.AddReceivers() instead of directly setting the serverURLs because we want to normalize the URLs.

// New returns a new instance of Bark service. You can use this service to send messages to bark. By default, the
// service will use the default server (https://api.day.app/).
func New(deviceKey string) *Service { _ = "STUB: not implemented"; return nil }

// postData is the data to send to the bark server.
type postData struct {
	DeviceKey string `json:"device_key"`
	Title     string `json:"title"`
	Body      string `json:"body,omitempty"`
	Badge     int    `json:"badge,omitempty"`
	Sound     string `json:"sound,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Group     string `json:"group,omitempty"`
	URL       string `json:"pushURL,omitempty"`
}

func (s *Service) send(ctx context.Context, serverURL, subject, content string) error {
	_ = "STUB: not implemented"
	return nil
}

// Marshal the message to post

// Create new request

// Send request

// Read response and verify success

// Send takes a message subject and a message content and sends them to bark application.
func (s *Service) Send(ctx context.Context, subject, content string) error {
	_ = "STUB: not implemented"
	return nil
}
