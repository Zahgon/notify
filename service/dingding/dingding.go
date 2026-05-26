package dingding

import (
	"context"

	"github.com/blinkbean/dingtalk"
)

// Service encapsulates the DingTalk client.
type Service struct {
	config Config
	client *dingtalk.DingTalk
}

// Config is the Service configuration.
type Config struct {
	Token  string
	Secret string
}

// New returns a new instance of a DingTalk notification service.
func New(cfg *Config) *Service { _ = "STUB: not implemented"; return nil }

// Send takes a message subject and a message content and sends them to all previously set users.
func (s *Service) Send(ctx context.Context, subject, content string) error {
	_ = "STUB: not implemented"
	return nil
}
