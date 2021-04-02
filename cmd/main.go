package main

import (
	"log"

	"github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/handler"
	"github.com/Yujiman/GoTodo/pkg/repository"
	"github.com/Yujiman/GoTodo/pkg/service"
)

func main() {

	rep := repository.NewRepository()
	service := service.NewService(rep)
	handler := handler.NewHandler(service)
	
	srv := new(TodoP.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal("Error running  http server: ", err.Error())
	}
}
