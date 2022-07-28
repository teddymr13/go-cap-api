package main

import (
	"capi/app"
	"capi/logger"
)

func main() {
	logger.Info("starting application...")
	app.Start()
}
