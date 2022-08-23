package handler

import (
	"io"
	"net/http"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "hello world")
}
