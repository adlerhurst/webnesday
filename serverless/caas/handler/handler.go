package handler

type Handler struct {
	w writer
	r reader
}

func New(w writer, r reader) *Handler {
	return &Handler{w, r}
}
