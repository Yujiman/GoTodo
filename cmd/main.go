package main

import (
	"log"

	TodoP "github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/handler"
	"github.com/Yujiman/GoTodo/pkg/repository"
	"github.com/Yujiman/GoTodo/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("errors init configs: %s", err.Error())
	}

	rep := repository.NewRepository()
	service := service.NewService(rep)
	handler := handler.NewHandler(service)

	srv := new(TodoP.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal("Error running  http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
