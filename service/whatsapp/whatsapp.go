package whatsapp

import (
	"context"
)

// Service encapsulates the WhatsApp client along with internal state for storing contacts.
type Service struct{}

// New returns a new instance of a WhatsApp notification service.
func New() (*Service, error) {
	_ = "STUB: not implemented"
	return nil,

		// LoginWithSessionCredentials provides helper for authentication using whatsapp.Session credentials.
		nil
}

func (s *Service) LoginWithSessionCredentials(_, _, _, _ string, _, _ []byte) error {
	_ = "STUB: not implemented"

	// LoginWithQRCode provides helper for authentication using QR code on terminal.
	// Refer: https://github.com/Rhymen/go-whatsapp#login for more information.
	return nil
}

func (s *Service) LoginWithQRCode() error {
	_ = "STUB: not implemented"

	// AddReceivers takes WhatsApp contacts and adds them to the internal contacts list. The Send method will send
	// a given message to all those contacts.
	return nil
}

func (s *Service) AddReceivers(_ ...string) {
	_ = "STUB: not implemented"

	// Send takes a message subject and a message body and sends them to all previously set contacts.
	return
}

func (s *Service) Send(_ context.Context, _, _ string) error { _ = "STUB: not implemented"; return nil }
