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
	cfg := config.New_config()

	db, err := repository.NewDB(*cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos)

	router := handlers.NewRouter(services)

	serv := server.NewApiServer(cfg, logrus.New(), router)

	if err := serv.Start(); err != nil {
		logrus.Fatal(err.Error())
	}

}
