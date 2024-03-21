package main

import (
	"github.com/brandaoplaster/gopportunities/config"
	"github.com/brandaoplaster/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	println("Starting...")

	logger = config.GetLogger("main")

	erro := config.Init()

	if erro != nil {
		logger.ErrorFormat("Config initialization failed: %v", erro)
		return
	}

	router.Initialize()
}
