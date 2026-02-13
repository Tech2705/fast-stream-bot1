package main

import (
	"github.com/biisal/fast-stream-bot/internal/console"
)

func main() {
	console.Info("This is an info message")
	console.Success("This is a success message")
	console.Warning("This is a warning message")
	console.Error("This is an error message")
	console.Debug("This is a debug message")
	// console.Fatal("This is a fatal message (should exit)")
}
