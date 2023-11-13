package main

import (
	"go-web-platform/config"
	"go-web-platform/logging"
	"go-web-platform/services"
)

func main() {
	services.RegisterDefaultServices()

	_, _ = services.Call(writeMessage)

	val := struct {
		message string
		logging.Logger
	}{
		message: "Hello from the struct",
	}
	_ = services.Populate(&val)
	val.Logger.Debug(val.message)
}

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section is not found")
	}
}
