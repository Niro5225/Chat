package server

import (
	"chat-app/internal/config"
	"chat-app/internal/handlers"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	conf   *config.Config
	logger *logrus.Logger
	router *handlers.Router
}

func NewApiServer(cfg *config.Config, log *logrus.Logger, router *handlers.Router) *ApiServer {
	return &ApiServer{
		conf:   cfg,
		logger: log,
		router: router,
	}
}

func (a *ApiServer) Start() error {
	//Настройка логера
	if err := a.configure_logger(); err != nil {
		return err
	}

	a.router.Configure_router()

	a.logger.Info("Starting API Server")

	return http.ListenAndServe(a.conf.Bind, a.router.Router)
}

func (a *ApiServer) configure_logger() error {
	level, err := logrus.ParseLevel(a.conf.Log_level)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)
	a.logger.SetFormatter(new(logrus.JSONFormatter))

	return nil
}
