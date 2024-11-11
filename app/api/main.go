package main

import (
	"gobby/src/logger"
	"gobby/src/routes"

	"os"
)

func main() {
	logger.InitLogger()
	routes.InitRoutes()
	// TODO defer logger.Sugar.Sync() and other graceful shutdown tasks

	// Tmp test code to show current working directory in container
	dir, err := os.Getwd()
	if err != nil {
		logger.Sugar.Fatal("an error occured while getting current working directory: %s", err)
	}
	logger.Sugar.Infof("running in %s", dir)
}
