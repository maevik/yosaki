package home

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

type Authenticator interface {
	IsAuthenticated(r *http.Request) bool
}

type Handler struct {
	renderer      *view.Renderer
	service       *Service
	authenticator Authenticator
}

// TODO

func NewHandler(renderer *view.Renderer, service *Service, authenticator Authenticator) *Handler {
	return &Handler{renderer: renderer, service: service, authenticator: authenticator}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if !h.authenticator.IsAuthenticated(r) {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	h.renderer.Render(w, "home", view.Data{Page: h.service.PageTitle()})
}
