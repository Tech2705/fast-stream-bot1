package main

import (
	"github.com/biisal/fast-stream-bot/config"
)

func main() {
	printLogo("v1.0.0")
	cfg := config.MustLoad("")
	if err := mount(cfg); err != nil {
		panic(err)
	}
}
