package main

func main() {
	logging := NewLogging()
	logging.Debugf("debug: %s, %d", "info", 1)
	logging.Infof("info: %s, %d", "info", 1)
	logging.Warnf("warn: %s, %d", "info", 1)
	logging.Errorf("error: %s, %d", "info", 1)

	logger := NewLogger()
	logger.Debugf("debug: %s, %d", "info", 1)
	logger.Infof("info: %s, %d", "info", 1)
	logger.Warnf("warn: %s, %d", "info", 1)
	logger.Errorf("error: %s, %d", "info", 1)
}
