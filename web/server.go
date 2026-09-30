package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/maevik/yosaki/web/auth"
	"github.com/maevik/yosaki/web/home"
	"github.com/maevik/yosaki/web/view"
)

func Router() http.Handler {
	templates := view.MustParseFiles("tmpl/**/*.tmpl")
	renderer := view.NewRenderer(templates)

	handler := NewHandler(renderer)

	authRepository := auth.NewRepository()
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(renderer, authService)

	homeRepository := home.NewRepository()
	homeService := home.NewService(homeRepository)
	homeHandler := home.NewHandler(renderer, homeService, authService)

	router := chi.NewRouter()

	router.Handle("/public/*",
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("web/public"))))

	router.Get("/", homeHandler.Index)
	router.Get("/signin", authHandler.Signin)
	router.Post("/signin", authHandler.SigninPost)
	router.Get("/signup", authHandler.Signup)
	router.Post("/signup", authHandler.SignupPost)

	router.NotFound(handler.NotFound)
	router.MethodNotAllowed(handler.MethodNotAllowed)

	return router
}
