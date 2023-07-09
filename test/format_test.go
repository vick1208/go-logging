package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("Hello Logging")
	log.Warn("Hello Logging")
	log.Error("Hello Logging")
}
