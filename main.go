package main

import "github.com/maevik/yosaki/web"

func main() {
	web := web.New()
	web.Register()
	web.Start()
}
