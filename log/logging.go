package main

import (
	"log"
	"os"
)

type Logging struct {
	logger *log.Logger
}

func NewLogging() *Logging {
	return &Logging{
		logger: log.New(os.Stderr, "", log.LstdFlags|log.LUTC|log.Lshortfile),
	}
}

func (l *Logging) Debugf(format string, v ...interface{}) {
	l.logger.SetPrefix("[DEBUG]\t")
	l.logger.Printf(format, v...)
}

func (l *Logging) Infof(format string, v ...interface{}) {
	l.logger.SetPrefix("[INFO]\t")
	l.logger.Printf(format, v...)
}

func (l *Logging) Warnf(format string, v ...interface{}) {
	l.logger.SetPrefix("[WARN]\t")
	l.logger.Printf(format, v...)
}

func (l *Logging) Errorf(format string, v ...interface{}) {
	l.logger.SetPrefix("[ERROR]\t")
	l.logger.Printf(format, v...)
}
