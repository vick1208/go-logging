package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	log := logrus.New()

	log.SetLevel(logrus.TraceLevel)

	log.Trace("Trace")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	// log.Fatal("Fatal")
}
