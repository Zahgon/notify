package syslog

import (
	"context"
	"io"
	"log/syslog"
)

// mockSyslogWriter abstracts log/syslog for writing unit tests.
type syslogWriter interface {
	io.WriteCloser
}

// Service encapsulates a syslog daemon writer.
type Service struct {
	writer syslogWriter
}

// dial is a wrapper function around syslog.Dial. It normalizes the prefix tag in case that it's empty and returns
// a new Service with the writer field set to writer from the call to syslog.Dial.
func dial(network, raddr string, priority syslog.Priority, tag string) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Usually we could call syslog.New and syslog.Dial respectively for specific use-cases. But since syslog.New is
// only a wrapper around a call to syslog.Dial without information about the network we're doing the same here to
// keep the API a little more clean.

// New returns a new instance of a Service notification service. Parameter 'tag' is used as a log prefix and may be left
// empty, it has a fallback value.
func New(priority syslog.Priority, tag string) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFromDial returns a new instance of a Service notification service. The underlying syslog writer establishes a
// connection to a log daemon by connecting to address raddr on the specified network. Parameter 'tag' is used as a log
// prefix and may be left empty, it has a fallback value.
// Calling NewFromDial with network and raddr being empty strings is equal in function to calling New directly.
func NewFromDial(network, raddr string, priority syslog.Priority, tag string) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close the underlying syslog writer.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

// Send takes a message subject and a message body and sends them to all previously set channels.
// user used for sending the message has to be a member of the channel.
func (s *Service) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}
