package cmd

import (
	"log"
	"net/http"

	"github.com/maevik/yosaki/web"
)

func Web() error {
	// TODO: graceful
	log.Printf("listening on http://localhost%s", ":8080")
	return http.ListenAndServe(":8080", web.Router())
}
