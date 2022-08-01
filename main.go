package main

import (
	"capi/app"
	"capi/logger"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TEST"))
	logger.Info("starting application...")
	app.Start()
}
