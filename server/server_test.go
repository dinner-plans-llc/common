package server_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/dinner-plans-llc/common/logging"
	"go.uber.org/fx"
)

func TestProvideServer(t *testing.T) {

	app := fx.New(
		fx.Provide(logging.ProvideLogger),
	)

	start, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	if err := app.Start(start); err != nil {
		log.Fatal(err)
	}

	stop, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	if err := app.Stop(stop); err != nil {
		log.Fatal(err)
	}
}
