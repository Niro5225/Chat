package apiserver

import (
	"chat/pkg/store"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Api_server struct {
	conf   *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	tpl    *template.Template
}

func New(conf *Config) Api_server {
	var err error
	var tpl *template.Template
	//Считывание всех gohtml файлов
	tpl, err = template.ParseGlob("pkg\\web\\*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	return Api_server{
		conf:   conf,
		logger: logrus.New(),    //Инициализация логера
		router: mux.NewRouter(), //Инициализация роутера
		tpl:    tpl,
	}
}

//Запуск сервера
func (a *Api_server) Start() error {
	//Настройка логера
	if err := a.configure_logger(); err != nil {
		return err
	}

	//Добавление Handle методов
	a.configure_router()

	//Настройка хранилища
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
	a.router.HandleFunc("/ping", a.handle_hello())
	a.router.HandleFunc("/", a.index())
	a.router.HandleFunc("/registration", a.registration())
	a.router.HandleFunc("/registration_process", a.RegistrationProcess())
	a.router.HandleFunc("/login", a.login())
	a.router.HandleFunc("/login_process", a.loginProcess())
	a.router.HandleFunc("/find_email", a.FindEmail())
	a.router.HandleFunc("/create_chat", a.CreateChat())
	a.router.HandleFunc("/chat_create_process", a.CreateChatProcess())
	a.router.HandleFunc("/open_chat", a.OpenChat())
	a.router.HandleFunc("/message_process", a.MessageProcess())
}

func (a *Api_server) configure_store() error {
	s := store.New(a.conf.Store)
	if err := s.Open(); err != nil {
		return err
	}

	a.store = s

	return nil
}
