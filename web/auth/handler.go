package auth

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

type Handler struct {
	renderer *view.Renderer
	service  *Service
}

func NewHandler(renderer *view.Renderer, service *Service) *Handler {
	return &Handler{renderer: renderer, service: service}
}

func (h *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "signin", view.Data{Page: "Sign In"})
}

func (h *Handler) SigninPost(w http.ResponseWriter, r *http.Request) {
	h.service.Signin(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "signup", view.Data{Page: "Sign Up"})
}

func (h *Handler) SignupPost(w http.ResponseWriter, r *http.Request) {
	h.service.Signup(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
