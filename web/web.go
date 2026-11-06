package web

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maevik/yosaki/web/auth"
	"github.com/maevik/yosaki/web/home"
	"github.com/maevik/yosaki/web/public"
	"github.com/maevik/yosaki/web/view"
)

type Web struct {
	router   *http.ServeMux
	server   *http.Server
	renderer *view.Renderer
}

func New() *Web {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	rdr := view.NewRenderer(
		view.ParseFiles("tmpl/**/*.tmpl"),
	)

	return &Web{
		router:   mux,
		server:   srv,
		renderer: rdr,
	}
}

func (w *Web) Register() {
	auth.Register(w.renderer, w.router)
	home.Register(w.renderer, w.router)
	public.Register(w.renderer, w.router)
}

func (w *Web) Start() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	go func() {
		slog.Info("start server", "addr", w.server.Addr)
		if err := w.server.ListenAndServe(); err != nil {
			slog.Error("failed to start server", "addr", w.server.Addr, "error", err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("shutdown server", "addr", w.server.Addr)
	if err := w.server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", "addr", w.server.Addr, "error", err)
	}
}
