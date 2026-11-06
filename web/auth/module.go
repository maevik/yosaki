package auth

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

func Register(rdr *view.Renderer, mux *http.ServeMux) {
	repo := newRepository()
	svc := newService(repo)
	h := newHandler(rdr, svc)

	mux.HandleFunc("GET /signin", h.Signin)
	mux.HandleFunc("POST /signin", h.SigninPost)
	mux.HandleFunc("GET /signup", h.Signup)
	mux.HandleFunc("POST /signup", h.SignupPost)
}
