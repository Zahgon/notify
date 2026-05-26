package googlechat

import (
	"context"

	"google.golang.org/api/chat/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type spacesMessageCreator interface {
	Create(string, *chat.Message) callCreator
}

type callCreator interface {
	Do(...googleapi.CallOption) (*chat.Message, error)
}

var (
	// Compile-time check to ensure that client implements the spaces message service
	// interface.
	_ spacesMessageCreator = new(messageCreator)
	// Compile-time check to ensure that client implements the create call interface.
	_ callCreator = new(chat.SpacesMessagesCreateCall)
)

// messageCreator is wrapper struct for the native chat.SpacesMessagesService struct.
// This exists so that we can mock the chat.SpacesMessagesCreateCall with the common
// interface "createCall".
type messageCreator struct {
	*chat.SpacesMessagesService
}

func newMessageCreator(ctx context.Context, options ...option.ClientOption) (spacesMessageCreator, error) {
	_ = "STUB: not implemented"
	return *new(spacesMessageCreator), nil
}

// Create creates a createCall struct for google chat. In order to execute sending
// the message utilize the `.Do` method found on the createCall.
func (m *messageCreator) Create(parent string, message *chat.Message) callCreator {
	_ = "STUB: not implemented"
	return *new(callCreator)
}

// Service encapsulates the google chat client along with internal state for storing
// chat spaces.
type Service struct {
	messageCreator spacesMessageCreator
	spaces         []string
}

// New returns an instance of the google chat notification service.
func New(options ...option.ClientOption) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWithContext returns an instance of the google chat notification service with the
// specified context. Utilize this constructor if the message requires the context to
// be set.
func NewWithContext(ctx context.Context, options ...option.ClientOption) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes a name of authorized spaces and appends them to the internal
// spaces slice. The Send method will send a given message to all those spaces.
func (s *Service) AddReceivers(spaces ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all the spaces
// previously set.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	// Treating subject as message title
	return nil
}
