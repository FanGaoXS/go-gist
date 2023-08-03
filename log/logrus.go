package main

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Entry
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetReportCaller(true)

	return &Logger{
		logger: logger.WithFields(logrus.Fields{
			"app": "go-gist",
			"env": "dev",
		}),
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
