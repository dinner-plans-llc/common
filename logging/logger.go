package logging

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

// ProviderLogger create new uber/zap logg instance
func ProvideLogger() (*zap.Logger, error) {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("zap.NewDevelopment err=%s", err)
	}

	logger.Debug("provide new logger")

	return logger, nil
}
