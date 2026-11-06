package home

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

type handler struct {
	renderer *view.Renderer
	service  *service
}

func newHandler(rdr *view.Renderer, svc *service) *handler {
	return &handler{
		renderer: rdr,
		service:  svc,
	}
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "home", view.Data{Page: h.service.PageTitle()})
}
