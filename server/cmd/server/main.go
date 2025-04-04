package main

import (
	"log"

	"github.com/brotigen23/goph-keeper/server/internal/app"
)

// @title My API
// @version 1.0
// @description goph-keeper
// @termsOfService http://example.com/terms/
// @contact.name API Support
func main() {
	log.Println("start server... ")
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
