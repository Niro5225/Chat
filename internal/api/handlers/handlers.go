package handlers

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/models"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router   *mux.Router
	services *chat_domain.Services
	tpl      *template.Template
}

func NewRouter(services *chat_domain.Services) *Router { //Создание роутера
	var tpl *template.Template

	tpl, err := template.ParseGlob("internal\\web\\*.gohtml") //получение всех gohtml файлов
	if err != nil {
		log.Fatalln(err)
	}
	return &Router{Router: mux.NewRouter(), services: services, tpl: tpl}
}

func (r *Router) Configure_router() { //Настройка роутера
	r.Router.HandleFunc("/ping", r.ping())
	r.Router.HandleFunc("/", r.test())

}

func (router *Router) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func (router *Router) test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := router.services.UserService.CreateUser(*models.NewUser("test", "test", "testemail"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		data := struct {
			fname string
			lname string
			em    string
		}{u.FirstName, u.LastName, u.Email}
		router.tpl.ExecuteTemplate(w, "index.gohtml", data)
	}
}
