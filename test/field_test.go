package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithField("username", "viddy").Info("Hello Warudo")
	log.WithField("username", "viddy").WithField("name", "Vicky Dien").Info("Hello Warudo")
}

func TestFields(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"username": "viddy",
		"name":     "Vicky Dien",
	}).Info("Hello World")

}
