package main

import (
	"gobby/src/handlers/healthcheck"
	"net/http"

	"os"

	"go.uber.org/zap"
)

func initApis() {
	http.HandleFunc("/", healthcheck.Healthcheck)
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	dir, err := os.Getwd()
	if err != nil {
		sugar.Fatal("an error occured while getting current working directory: %s", err)
	}

	initApis()
	sugar.Infof("running in %s", dir)
	sugar.Infof("api listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
