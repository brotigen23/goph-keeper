package main

import (
	"flag"

	"github.com/brotigen23/goph-keeper/client/internal/cli/app"
)

func main() {
	isTUI := flag.Bool("t", false, "Run in tui")
	flag.Parse()

	if *isTUI {
	} else {
		app := app.New()
		app.Run()
	}
}
