package handler

import (
	"context"
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed result.html
	result         string
	resultTemplate = template.Must(template.New("result").Parse(result))
)

type ResultData struct {
	Attended string
	Count    int
}

type reader interface {
	Get(ctx context.Context) ([]*ResultData, error)
}

func (h *Handler) HandleResult(w http.ResponseWriter, r *http.Request) {
	res, err := h.r.Get(r.Context())
	if err != nil {
		log.Println("unable to get", err)
	}

	logOnErr("error durring result rendering:", resultTemplate.Execute(w, res))
}

func logOnErr(cause string, err error) {
	if err != nil {
		log.Println(cause, err)
	}
}
