package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"go.uber.org/zap"
)

// App
type App struct {
	app    *firebase.App
	logger *zap.SugaredLogger
}

// Auth
func (a *App) Auth(ctx context.Context) (*auth.Client, error) {
	return a.app.Auth(ctx)
}

// Messaging
func (a *App) Messaging(ctx context.Context) (*messaging.Client, error) {
	return a.app.Messaging(ctx)
}
