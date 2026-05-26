package matrix

import (
	"context"

	matrix "maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type matrixClient interface {
	SendMessageEvent(
		ctx context.Context,
		roomID id.RoomID,
		eventType event.Type,
		contentJSON any,
		extra ...matrix.ReqSendEvent,
	) (resp *matrix.RespSendEvent, err error)
}

// Compile time check to ensure that matrix.Client implements the matrixClient interface.
var _ matrixClient = new(matrix.Client)

// New returns a new instance of a Matrix notification service.
// For more information about the Matrix api specs:
//
// -> https://spec.matrix.org/v1.2/client-server-api
func New(userID id.UserID, roomID id.RoomID, homeServer, accessToken string) (*Matrix, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send takes a message body and sends them to the previously set channel.
// you will need an account, access token and roomID
// see https://matrix.org
func (s *Matrix) Send(ctx context.Context, _, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func createMessage(message string) Message { _ = "STUB: not implemented"; return *new(Message) }
