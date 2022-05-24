package logging_test

import (
	"testing"

	"github.com/dinner-plans-llc/common/logging"
)

func TestProvideLogger(t *testing.T) {

	logger, err := logging.ProvideLogger()
	if err != nil {
		t.Error(err)
	}

	logger.Debug("logger successfully created")
}
