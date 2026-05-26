package mailgun

// Option describes a functional parameter for the Mailgun constructor.
type Option func(*Mailgun)

// WithEurope sets the API Mailgun base url to Europe region.
func WithEurope() Option { _ = "STUB: not implemented"; return *new(Option) }
