package main

import (
	"github.com/brandaoplaster/gopportunities/config"
	"github.com/brandaoplaster/gopportunities/router"
)

func main() {
	println("Starting...")

	erro := config.Init()

	if erro != nil {
		println("Error initializing config")
		return
	}

	router.Initialize()
}
