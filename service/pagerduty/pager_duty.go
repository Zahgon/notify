package pagerduty

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"

	"github.com/nikoksr/notify"
)

type Client interface {
	CreateIncidentWithContext(
		ctx context.Context,
		from string,
		options *pagerduty.CreateIncidentOptions,
	) (*pagerduty.Incident, error)
}

// Compile-time check to verify that the PagerDuty type implements the notifier.Notifier interface.
var _ notify.Notifier = &PagerDuty{}

type PagerDuty struct {
	*Config

	Client Client
}

func New(token string, clientOptions ...pagerduty.ClientOptions) (*PagerDuty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PagerDuty) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// set the service ID to the receiver

func (s *PagerDuty) IncidentOptions(subject, message string) *pagerduty.CreateIncidentOptions {
	_ = "STUB: not implemented"
	return nil
}

// service ID will be set per receiver
