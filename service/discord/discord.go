package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

//go:generate mockery --name=discordSession --output=. --case=underscore --inpackage
type discordSession interface {
	ChannelMessageSend(channelID string, content string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

// Compile-time check to ensure that discordgo.Session implements the discordSession interface.
var _ discordSession = new(discordgo.Session)

// Discord struct holds necessary data to communicate with the Discord API.
type Discord struct {
	client     discordSession
	channelIDs []string
}

// New returns a new instance of a Discord notification service.
// The instance is created with a default Discord session. For advanced configuration
// such as proxy support or custom timeouts, use SetClient to provide a custom session
// before calling the authentication methods.
func New() *Discord { _ = "STUB: not implemented"; return nil }

// authenticate will configure authentication on the existing Discord session.
// If a custom session was set via SetClient, it preserves all custom configuration
// (HTTP client, proxy settings, etc.) while setting the authentication token.
func (d *Discord) authenticate(token string) error {
	_ = "STUB: not implemented"
	// If we have an existing session, configure it in-place
	return nil
}

// Fallback: create new session (only if client is nil or not a Session)

// DefaultSession returns a new Discord session with default configuration.
// This allows configuring a custom session without importing the discordgo package.
// The session is created with Discord's standard defaults and can be customized before
// passing it to SetClient.
//
// Example - Configure proxy without importing discordgo:
//
//	session := discord.DefaultSession()
//	session.Client = &http.Client{
//	    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
//	}
//	d := discord.New()
//	d.SetClient(session)
//	d.AuthenticateWithBotToken(token)
func DefaultSession() *discordgo.Session {
	_ = "STUB: not implemented"
	// Use discordgo.New with empty token to get proper defaults
	return nil
}

// SetClient sets a custom Discord session, allowing full control over the
// session configuration including HTTP client, proxy settings, intents, and
// other Discord session options.
//
// This is useful for advanced use cases like proxy configuration or custom timeouts.
// After calling SetClient, you can still use the AuthenticateWith* methods to set
// the authentication token while preserving your custom configuration.
//
// Use DefaultSession() to get a session instance without importing discordgo.
//
// Example - Configure proxy, then authenticate:
//
//	session := discord.DefaultSession()
//	session.Client = &http.Client{
//	    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
//	}
//	d := discord.New()
//	d.SetClient(session)
//	d.AuthenticateWithBotToken(token)  // Preserves proxy configuration
func (d *Discord) SetClient(client *discordgo.Session) {
	_ = "STUB: not implemented"

	// AuthenticateWithBotToken authenticates you as a bot to Discord via the given access token.
	// For more info, see here: https://pkg.go.dev/github.com/bwmarrin/discordgo@v0.22.1#New
	return
}

func (d *Discord) AuthenticateWithBotToken(token string) error {
	_ = "STUB: not implemented"
	return nil
}

// AuthenticateWithOAuth2Token authenticates you to Discord via the given OAUTH2 token.
// For more info, see here: https://pkg.go.dev/github.com/bwmarrin/discordgo@v0.22.1#New
func (d *Discord) AuthenticateWithOAuth2Token(token string) error {
	_ = "STUB: not implemented"
	return nil
}

// parseBotToken parses a regular token to a bot token that is understandable for discord.
// For more info, see here: https://pkg.go.dev/github.com/bwmarrin/discordgo@v0.22.1#New
func parseBotToken(token string) string { _ = "STUB: not implemented"; return "" }

// parseBotToken parses a regular token to a OAUTH2 token that is understandable for discord.
// For more info, see here: https://pkg.go.dev/github.com/bwmarrin/discordgo@v0.22.1#New
func parseOAuthToken(token string) string { _ = "STUB: not implemented"; return "" }

// AddReceivers takes Discord channel IDs and adds them to the internal channel ID list. The Send method will send
// a given message to all those channels.
func (d *Discord) AddReceivers(channelIDs ...string) { _ = "STUB: not implemented"; return }

// Send takes a message subject and a message body and sends them to all previously set chats.
func (d Discord) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Treating subject as message title
