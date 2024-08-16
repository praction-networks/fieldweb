package main

import (
	"fieldweb/src/logger"
	"fieldweb/src/startup"
	"log"
)

func main() {

	if err := startup.FieldWebStart(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	// Example log messages
	logger.Info("FieldWeb Application started", "version", "1.0.0")

}
