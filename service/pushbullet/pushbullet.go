package pushbullet

import (
	"context"

	"github.com/cschomburg/go-pushbullet"
)

// Pushbullet struct holds necessary data to communicate with the Pushbullet API.
type Pushbullet struct {
	client          *pushbullet.Client
	deviceNicknames []string
}

// New returns a new instance of a Pushbullet notification service.
// For more information about Pushbullet api token:
//
//	-> https://docs.pushbullet.com/#api-overview
func New(apiToken string) *Pushbullet { _ = "STUB: not implemented"; return nil }

// AddReceivers takes Pushbullet device nicknames and adds them to the internal deviceNicknames list.
// The Send method will send a given message to all those devices.
func (pb *Pushbullet) AddReceivers(deviceNicknames ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all valid devices.
// you will need Pushbullet installed on the relevant devices
// (android, chrome, firefox, windows)
// see https://www.pushbullet.com/apps
func (pb Pushbullet) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
