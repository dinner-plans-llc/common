package notification

import "context"

// INotification notification methods definitions
type INotification interface {
	SendDevice(ctx context.Context, token, title, body string) (string, error)
	SendTopic(ctx context.Context, topic, title, body string) (string, error)
	Subscribe(ctx context.Context, topic, token string) error
	UnSubscribe(ctx context.Context, topic, token string) error
}
