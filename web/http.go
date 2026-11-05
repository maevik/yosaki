package web

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WebServer struct {
	Router *http.ServeMux
	Server *http.Server
}

func NewServer() *WebServer {
	r := http.NewServeMux()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return &WebServer{
		Router: r,
		Server: s,
	}
}

func (ws *WebServer) RunServer() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	go func() {
		slog.Info("start web server", "addr", ws.Server.Addr)

		if err := ws.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start web server", "addr", ws.Server.Addr, "err", err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("shutdown web server", "addr", ws.Server.Addr)

	if err := ws.Server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown web server", "addr", ws.Server.Addr, "err", err)
	}
}
