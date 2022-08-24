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

	for _, r := range res {
		r.Attended = translateAttended(r.Attended)
	}

	logOnErr("error durring result rendering:", resultTemplate.Execute(w, res))
}

func logOnErr(cause string, err error) {
	if err != nil {
		log.Println(cause, err)
	}
}

func translateAttended(attended string) string {
	switch attended {
	case "first":
		return "First time"
	case "low":
		return "2 to 5"
	case "medium":
		return "6 to 10"
	case "high":
		return "more than 10"
	case "all":
		return "All (Pascal)"
	default:
		return ""
	}
}
