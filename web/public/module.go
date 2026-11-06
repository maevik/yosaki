package public

import (
	"net/http"

	"github.com/maevik/yosaki/web/view"
)

func Register(rdr *view.Renderer, mux *http.ServeMux) {
	mux.Handle("GET /public/{path...}",
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("web/public"))))

	mux.HandleFunc("/", rdr.RenderError)
}
