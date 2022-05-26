package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

// Config
type Config struct {
	AppName string
	FBCfg   *firebase.Config
	Opts    option.ClientOption
}

// ProvideApp
func (c *Config) ProvideApp(ctx context.Context, log *zap.Logger) (*App, error) {

	var app App
	var err error

	app.logger = log.Named(c.AppName).Sugar()

	app.app, err = firebase.NewApp(ctx, c.FBCfg, c.Opts)
	if err != nil {
		app.logger.Error("failed to initialize firebase app ", err)
		return nil, err
	}

	return &app, nil
}
