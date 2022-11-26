package handlers

import (
	"chat-app/internal/service"
	"chat-app/models"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router   *mux.Router
	services *service.Services
	tpl      *template.Template
}

func NewRouter(services *service.Services) *Router {
	var tpl *template.Template
	//Считывание всех gohtml файлов
	tpl, err := template.ParseGlob("internal\\web\\*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	return &Router{Router: mux.NewRouter(), services: services, tpl: tpl}
}

func (r *Router) Configure_router() {
	r.Router.HandleFunc("/ping", r.ping())
	r.Router.HandleFunc("/test", r.test())
}

func (router *Router) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func (router *Router) test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := router.services.UserRepository.CreateUser(*models.NewUser("test", "test", "testemail"))
		if err != nil {
			io.WriteString(w, err.Error())
		}

		userCredential := models.NewUserCredential(u.ID, "testpassword", u.Email)

		userCredential, err = router.services.UserRepository.CreateUserCredential(*userCredential)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	}
}
