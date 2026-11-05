package web

import (
	"net/http"

	"github.com/maevik/yosaki/web/auth"
	"github.com/maevik/yosaki/web/home"
	"github.com/maevik/yosaki/web/view"
)

func (ws *WebServer) NewRouter() {
	templates := view.MustParseFiles("tmpl/**/*.tmpl")
	renderer := view.NewRenderer(templates)

	authService := auth.NewService(auth.NewRepository())
	authHandler := auth.NewHandler(renderer, authService)

	homeService := home.NewService(home.NewRepository())
	homeHandler := home.NewHandler(renderer, homeService, authService)

	ws.Router.Handle("GET /public/{path...}",
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("web/public"))))

	ws.Router.HandleFunc("GET /", homeHandler.Index)

	ws.Router.HandleFunc("GET /signin", authHandler.Signin)
	ws.Router.HandleFunc("POST /signin", authHandler.SigninPost)
	ws.Router.HandleFunc("GET /signup", authHandler.Signup)
	ws.Router.HandleFunc("POST /signup", authHandler.SignupPost)

	ws.Router.HandleFunc("/", renderer.RenderError)
}
