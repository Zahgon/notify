package wechat

import (
	"context"
	"net/http"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

const defaultTimeout = 20 * time.Second

type verificationCallbackFunc func(r *http.Request, verified bool)

// Config is the Service configuration.
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Cache          cache.Cache
}

// wechatMessageManager abstracts go-wechat's message.Manager for writing unit tests.
type wechatMessageManager interface {
	Send(msg *message.CustomerMessage) error
}

// Service encapsulates the WeChat client along with internal state for storing users.
type Service struct {
	config         *Config
	messageManager wechatMessageManager
	userIDs        []string
}

// New returns a new instance of a WeChat notification service.
func New(cfg *Config) *Service { _ = "STUB: not implemented"; return nil }

// waitForOneOffVerification waits for the verification call from the WeChat backend.
//
// Should be running when (re-)applying settings in wechat configuration.
//
// Set devMode to true when using the sandbox.
//
// See https://developers.weixin.qq.com/doc/offiaccount/en/Basic_Information/Access_Overview.html
func (s *Service) waitForOneOffVerification(
	server *http.Server,
	devMode bool,
	callback verificationCallbackFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

// verification done; dev mode

// perform signature check

// verification done; prod mode

// verification not done (keep waiting)

// wait until verification is done and shutdown the server

// WaitForOneOffVerificationWithServer allows you to use WaitForOneOffVerification with a fully custom HTTP server.
//
// Should be running when (re-)applying settings in wechat configuration.
//
// Set devMode to true when using the sandbox.
//
// See https://developers.weixin.qq.com/doc/offiaccount/en/Basic_Information/Access_Overview.html
func (s *Service) WaitForOneOffVerificationWithServer(
	server *http.Server,
	devMode bool,
	callback verificationCallbackFunc,
) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForOneOffVerification waits for the verification call from the WeChat backend. It uses an internal
// ReadHeaderTimeout of 20 seconds to avoid blocking the caller for too long (potential slow loris attack). In case
// that you want to use a different timeout, you can use the WaitForOneOffVerificationWithServer method instead. It
// allows you to specify a custom server.
//
// Should be running when (re-)applying settings in wechat configuration.
//
// Set devMode to true when using the sandbox.
//
// See https://developers.weixin.qq.com/doc/offiaccount/en/Basic_Information/Access_Overview.html
func (s *Service) WaitForOneOffVerification(serverURL string, devMode bool, callback verificationCallbackFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// AddReceivers takes user ids and adds them to the internal users list. The Send method will send
// a given message to all those users.
func (s *Service) AddReceivers(userIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message content and sends them to all previously set users.
func (s *Service) Send(ctx context.Context, subject, content string) error {
	_ = "STUB: not implemented"
	return nil
}
