package main

import (
	"os"

	TodoP "github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/handler"
	"github.com/Yujiman/GoTodo/pkg/repository"
	"github.com/Yujiman/GoTodo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("errors init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable:%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialization DB: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	handler := handler.NewHandler(service)

	srv := new(TodoP.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal("Error running  http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
