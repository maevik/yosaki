package cmd

import "github.com/maevik/yosaki/web"

func RunWeb() {
	ws := web.NewServer()
	ws.NewRouter()
	ws.RunServer()
}
