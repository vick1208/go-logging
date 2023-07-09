package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello User")
	logrus.Warn("Hello User")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello User")
	logrus.Warn("Hello User")

}
