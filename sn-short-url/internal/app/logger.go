package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func setLogger(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logLevel)
	}

	logrus.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat: "2004-01-01 15:04:01",
		},
	)

	logrus.SetOutput(os.Stdout)
}
