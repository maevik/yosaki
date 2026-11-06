package auth

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

func (h *handler) Signin(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "signin", view.Data{Page: "Sign In"})
}

func (h *handler) SigninPost(w http.ResponseWriter, r *http.Request) {
	h.service.Signin(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *handler) Signup(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "signup", view.Data{Page: "Sign Up"})
}

func (h *handler) SignupPost(w http.ResponseWriter, r *http.Request) {
	h.service.Signup(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
