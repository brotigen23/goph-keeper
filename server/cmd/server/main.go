package main

import (
	"log"

	"github.com/brotigen23/goph-keeper/server/internal/app"
)

// @title GophKeeper API
// @version 1.0
// @description Сервис для хранения и менеджмента пользовательских данных

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Введите "Bearer <JWT>", где JWT - ваш access токен

// @host localhost:8080
// @BasePath /
func main() {
	log.Println("start server... ")
	err := app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
