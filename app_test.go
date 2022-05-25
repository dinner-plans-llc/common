package app_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/dinner-plans-llc/common/logging"
	"go.uber.org/fx"
)

func TestApp(t *testing.T) {

	app := fx.New(
		fx.Provide(logging.ProvideLogger),
		fx.Provide(ProvideHttpHandler),
		fx.Provide(ProvideHttpServer),
	)

	start, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Start(start); err != nil {
		t.Error(err)
	}

	stop, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Stop(stop); err != nil {
		t.Error(err)
	}
}

// ProvideHttpHandler
func ProvideHttpHandler() http.Handler {
	return http.NewServeMux()
}

// ProvideHttpServer
func ProvideHttpServer(handler http.Handler) *http.Server {
	s := &http.Server{
		Handler: handler,
	}

	return s
}
