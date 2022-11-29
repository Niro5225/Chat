package main

import (
	"chat-app/internal/config"
	"chat-app/internal/handlers"
	"chat-app/internal/repository"
	"chat-app/internal/server"
	"chat-app/internal/service"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New_config() //обьект конфига

	db, err := repository.NewDB(*cfg.DB) //обьект БД
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)  //обьект репозитория
	services := service.NewServices(repos) //обьект сервиса

	router := handlers.NewRouter(services)                 //обьект роутера
	serv := server.NewApiServer(cfg, logrus.New(), router) //обьект сервера

	if err := serv.Start(); err != nil { //запуск сервера
		logrus.Fatal(err.Error())
	}

}
