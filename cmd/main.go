package main

import (
	"github.com/myatakolya/API"
	"github.com/myatakolya/API/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(API.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatal("Ошибка при запуске сервера http server: %s", err.Error())
	}
}
