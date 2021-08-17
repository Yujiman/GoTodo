package main

import (
	"os"
	"os/signal"
	"syscall"

	todo "github.com/Yujiman/GoTodo"
	"github.com/Yujiman/GoTodo/pkg/handler"
	"github.com/Yujiman/GoTodo/pkg/repository"
	"github.com/Yujiman/GoTodo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	srv := new(todo.Server)
	//Graceful Shutdown
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatal("Error running  http server: ", err.Error())
		}
	}()
	logrus.Println("Todo started")

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGTERM, syscall.SIGINT)
	<-quite

	logrus.Println("Todo shutting down ")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")

	return viper.ReadInConfig()
}
