package handler

import (
	"context"
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed form.html
	form         string
	formTemplate = template.Must(template.New("form").Parse(form))
)

type formData struct {
	Err string
}

type writer interface {
	Save(ctx context.Context, attended string) error
}

func (h *Handler) HandleForm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleFormPOST(w, r)
	default:
		h.handleFormGET(w, r)
	}
}

func (h *Handler) handleFormGET(w http.ResponseWriter, r *http.Request) {
	formTemplate.Execute(w, new(formData))
}

type request struct {
	Attended string
}

func (h *Handler) handleFormPOST(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorOccured(w, err, "unable to parse form")
		return
	}

	if err := h.w.Save(r.Context(), r.Form.Get("attended")); err != nil {
		errorOccured(w, err, "unable to save form")
		return
	}

	http.Redirect(w, r, "result", http.StatusFound)
}

func errorOccured(w http.ResponseWriter, err error, text string) {
	log.Printf("%s: %v", text, err)
	w.WriteHeader(http.StatusNotAcceptable)
	formTemplate.Execute(w, formData{Err: "unable to parse form"})
}
