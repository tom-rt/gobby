package main

import (
	"fmt"
	"net/http"

	"os"

	"go.uber.org/zap"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World !")
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	dir, err := os.Getwd()
	if err != nil {
		sugar.Fatal("an error occured while getting current working directory: %s", err)
	}

	http.HandleFunc("/", helloWord)
	sugar.Infof("running in %s", dir)
	sugar.Infof("api listening on port 9999")
	http.ListenAndServe(":9999", nil)
}
