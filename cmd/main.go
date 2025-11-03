package main

import (
	"github.com/martyushova/posts/internal/app"
)

func main() {
	application := app.NewApp()
	application.Run()
}
