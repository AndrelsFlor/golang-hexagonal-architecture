package main

import (
	"rest_api/app"
	"rest_api/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
