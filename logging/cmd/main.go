package main

import (
	logging "logging/logging/internal/service"
)

func main() {
	loggingService := logging.NewLoggingService()
	loggingService.Run()
}
