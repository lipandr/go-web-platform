package main

import (
	"go-web-platform/logging"
)

func main() {
	var logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}

func writeMessage(log logging.Logger) {
	log.Trace("Trace")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warning")
	log.Panic("Panic")
}
