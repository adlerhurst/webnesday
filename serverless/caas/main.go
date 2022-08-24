package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/adlerhurst/webnesday/serverless/caas/handler"
	"github.com/adlerhurst/webnesday/serverless/caas/storage"
)

var (
	port = "8080"
)

func main() {
	overwriteEnv()
	h := handler.New(storage.NewPubsubWriter(), storage.NewCRDBReader())
	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/gui/form", h.HandleForm)
	http.HandleFunc("/gui/result", h.HandleResult)
	log.Fatal(http.ListenAndServe(port, nil))
}

func overwriteEnv() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
}
