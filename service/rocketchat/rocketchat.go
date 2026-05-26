package rocketchat

import (
	"context"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

// RocketChat struct holds necessary data to communicate with the RocketChat API.
type RocketChat struct {
	client       *rest.Client
	channelNames []string
}

// New returns a new instance of a RocketChat notification service.
// serverURL is the endpoint of server i.e "localhost" , scheme is protocol i.e "http/https"
// userID and token of the user sending the message.
func New(serverURL, scheme, userID, token string) (*RocketChat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes rocketchat channel names and adds them to the internal channel list. The Send method will send
// a given message to all channels in the list.
func (r *RocketChat) AddReceivers(channelNames ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set channels.
// user used for sending the message has to be a member of the channel.
// https://docs.rocket.chat/api/rest-api/methods/chat/postmessage
func (r *RocketChat) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
