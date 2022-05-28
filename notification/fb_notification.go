package notification

import (
	"context"

	"firebase.google.com/go/v4/messaging"
)

// FBV4 firebase v4 messaging implementation of INotification
type FBV4 struct {
	c *messaging.Client
}

// NewMgr create new firebase v4 client notification instance
func NewMgr(client *messaging.Client) *FBV4 {
	return &FBV4{c: client}
}

// SendDevice publish notification to a single device
func (fb *FBV4) SendDevice(ctx context.Context, token, title, body string) (string, error) {

	m := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	result, err := fb.c.Send(ctx, m)
	if err != nil {
		return "", err
	}

	return result, nil
}

// SendTopic publish notification a token
func (fb *FBV4) SendTopic(ctx context.Context, topic, title, body string) (string, error) {

	m := &messaging.Message{
		Topic: topic,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	result, err := fb.c.Send(ctx, m)
	if err != nil {
		return "", err
	}

	return result, nil
}

// Subscribe subscribe a single device to a topic
func (fb *FBV4) Subscribe(ctx context.Context, topic, token string) error {

	_, err := fb.c.SubscribeToTopic(ctx, []string{token}, topic)
	if err != nil {
		return err
	}

	return nil
}

// UnSubscribe unsubscribe a single device from a topic
func (fb *FBV4) UnSubscribe(ctx context.Context, topic, token string) error {

	_, err := fb.c.UnsubscribeFromTopic(ctx, []string{token}, topic)
	if err != nil {
		return err
	}

	return nil
}
