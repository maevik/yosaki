package web

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

type Handler struct {
	renderer *view.Renderer
}

func NewHandler(r *view.Renderer) *Handler {
	return &Handler{renderer: r}
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderStatus(w, http.StatusNotFound, "error", view.Data{
		Page:    "Not Found",
		Code:    http.StatusNotFound,
		Message: "The page you are trying to reach does not exist.",
	})
}

func (h *Handler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	h.renderer.RenderStatus(w, http.StatusMethodNotAllowed, "error", view.Data{
		Page:    "Method Not Allowed",
		Code:    http.StatusMethodNotAllowed,
		Message: "The page you are trying to reach does allow this method.",
	})
}
