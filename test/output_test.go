package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	log := logrus.New()

	file, _ := os.OpenFile("../application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.Info("Hello Log")
	log.Warn("Hello Log")
	log.Error("Hello Log")

}
