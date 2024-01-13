package logger

import (
	"gobby/src/handlers/healthcheck"
	"net/http"

	"os"

	"go.uber.org/zap"
)

Logger, _ := zap.NewProduction()

func initLogger() {
	defer logger.Sync()
	sugar := logger.Sugar()

	dir, err := os.Getwd()
	if err != nil {
		sugar.Fatal("an error occured while getting current working directory: %s", err)
	}

	http.HandleFunc("/", healthcheck.Healthcheck)
	sugar.Infof("running in %s", dir)
	sugar.Infof("api listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
