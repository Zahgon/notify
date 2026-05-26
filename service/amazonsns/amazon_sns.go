package amazonsns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// snsSendMessageAPI Basic interface to send messages through SNS.
//
//go:generate mockery --name=snsSendMessageAPI --output=. --case=underscore --inpackage
type snsSendMessageAPI interface {
	SendMessage(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

// snsSendMessageClient Client specific for SNS using aws sdk v2.
type snsSendMessageClient struct {
	client *sns.Client
}

// SendMessage Client specific for SNS using aws sdk v2.
func (s snsSendMessageClient) SendMessage(ctx context.Context,
	params *sns.PublishInput,
	optFns ...func(*sns.Options),
) (*sns.PublishOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AmazonSNS Basic structure with SNS information.
type AmazonSNS struct {
	sendMessageClient snsSendMessageAPI
	queueTopics       []string
}

// New creates a new AmazonSNS.
func New(accessKeyID, secretKey, region string) (*AmazonSNS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddReceivers takes queue urls and adds them to the internal topics
// list. The Send method will send a given message to all those
// Topics.
func (s *AmazonSNS) AddReceivers(queues ...string) { _ = "STUB: not implemented"; return }

// Send message to everyone on all topics.
func (s AmazonSNS) Send(ctx context.Context, subject, message string) error {
	_ = "STUB: not implemented"
	// For each topic
	return nil
}

// Create new input with subject, message and the specific topic

// Send the message
