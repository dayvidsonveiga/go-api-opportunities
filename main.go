package main

import (
	"github.com/dayvidsonveiga/go-api-opportunities/config"
	"github.com/dayvidsonveiga/go-api-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
