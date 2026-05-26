package pushbullet

import (
	"context"

	"github.com/cschomburg/go-pushbullet"
)

// SMS struct holds necessary data to communicate with the Pushbullet SMS API.
type SMS struct {
	client           *pushbullet.Client
	deviceIdentifier string
	phoneNumbers     []string
}

// NewSMS returns a new instance of a SMS notification service
// tied to an SMS capable device. deviceNickname is the
// Pushbullet nickname of the sms capable device from which messages are sent.
// (https://help.pushbullet.com/articles/how-do-i-send-text-messages-from-my-computer/).
// For more information about Pushbullet api token:
//
//	-> https://docs.pushbullet.com/#api-overview
func NewSMS(apiToken, deviceNickname string) (*SMS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes phone numbers and adds them to the internal phoneNumbers list. The Send method will send
// a given message to all registered phone numbers.
func (sms *SMS) AddReceivers(phoneNumbers ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all phone numbers.
// see https://help.pushbullet.com/articles/how-do-i-send-text-messages-from-my-computer/
func (sms SMS) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
