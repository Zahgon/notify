package pagerduty

import (
	"github.com/PagerDuty/go-pagerduty"
)

const (
	APIReferenceType        = "service_reference"
	APIPriorityReference    = "priority_reference"
	DefaultNotificationType = "incident"
)

// Config contains the configuration for the PagerDuty service.
type Config struct {
	FromAddress      string
	Receivers        []string
	NotificationType string
	Urgency          string
	PriorityID       string
}

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// OK checks if the configuration is valid.
// It returns an error if the configuration is invalid.
func (c *Config) OK() error { _ = "STUB: not implemented"; return nil }

// PriorityReference returns the PriorityID reference if it is set, otherwise it returns nil.
func (c *Config) PriorityReference() *pagerduty.APIReference { _ = "STUB: not implemented"; return nil }

// SetFromAddress sets the from address in the configuration.
func (c *Config) SetFromAddress(fromAddress string) { _ = "STUB: not implemented"; return }

// AddReceivers appends the receivers to the configuration.
func (c *Config) AddReceivers(receivers ...string) { _ = "STUB: not implemented"; return }

// SetPriorityID sets the PriorityID in the configuration.
func (c *Config) SetPriorityID(priorityID string) { _ = "STUB: not implemented"; return }

// SetUrgency sets the urgency in the configuration.
func (c *Config) SetUrgency(urgency string) { _ = "STUB: not implemented"; return }

// SetNotificationType sets the notification type in the configuration.
// If the notification type is empty, it will be set to the default value "incident".
func (c *Config) SetNotificationType(notificationType string) { _ = "STUB: not implemented"; return }
