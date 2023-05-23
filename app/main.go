package main

import (
	"net/http"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func helloGobby(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Gobby !"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		sugar.Errorf("an error occured while writing response's body content: %s", err)
	} else {
		sugar.Info("request treated successfuly")
	}
}

func initEndpoints() {
	http.HandleFunc("/", helloGobby)
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar = logger.Sugar()

	sugar.Info("initiating endpoints...")
	initEndpoints()
	sugar.Info("endpoints initiated successuly")

	sugar.Info("listening on port 1804...")
	http.ListenAndServe(":1804", nil)
}
