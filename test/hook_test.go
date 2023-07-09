package test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampHook struct {
}

func (s *SampHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampHook) Fire(ent *logrus.Entry) error {
	fmt.Println("Sample Hook", ent.Level, ent.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Error")
	logger.Debug("Hello Debug")

}
