package main

import (
	"log"

	"github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(TodoP.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal("Error running  http server: ", err.Error())
	}
}
