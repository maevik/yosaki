package home

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

func Register(rdr *view.Renderer, mux *http.ServeMux) {
	repo := newRepository()
	svc := newService(repo)
	h := newHandler(rdr, svc)

	mux.HandleFunc("GET /", h.Index)
}
