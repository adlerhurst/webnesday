package handler

import (
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/gui/form", http.StatusFound)
}
