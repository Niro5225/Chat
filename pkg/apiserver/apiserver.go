package apiserver

import (
	"chat/pkg/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api_server struct {
	conf   *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(conf *Config) Api_server {
	return Api_server{
		conf:   conf,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *Api_server) Start() error {
	if err := a.configure_logger(); err != nil {
		return err
	}

	a.configure_router()

	if err := a.configure_store(); err != nil {
		return err
	}

	a.logger.Info("Starting API Server")

	return http.ListenAndServe(a.conf.bind, a.router)
}

func (a *Api_server) configure_logger() error {
	level, err := logrus.ParseLevel(a.conf.log_level)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}

func (a *Api_server) configure_router() {
	a.router.HandleFunc("/hello", a.handle_hello())
}

func (a *Api_server) configure_store() error {
	s := store.New(a.conf.Store)
	if err := s.Open(); err != nil {
		return err
	}

	a.store = s

	return nil
}
