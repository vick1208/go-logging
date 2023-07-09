package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("Hello Log")
	logger.WithField("username", "eko").Info("Hello Log")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "soegianto")
	entry.Info("Hello Entry")
}
