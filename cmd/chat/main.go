package main

import (
	"chat-app/internal/api/handlers"
	"chat-app/internal/api/server"
	"chat-app/internal/chat/chat_database"
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New_config() //обьект конфига

	db, err := chat_database.NewDB(*cfg) //обьекconfig
	if err != nil {
		logrus.Fatal(err)
	}

	repos := chat_domain.NewRepository(db)     //обьект репозитория
	services := chat_domain.NewServices(repos) //обьект сервиса

	router := handlers.NewRouter(services)                 //обьект роутера
	serv := server.NewApiServer(cfg, logrus.New(), router) //обьект сервера

	if err := serv.Start(); err != nil { //запуск сервера
		logrus.Fatal(err.Error())
	}

}
