package handlers

import (
	"chat-app/internal/service"
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

func NewRouter(services *service.Services) *Router { //Создание роутера
	var tpl *template.Template

	tpl, err := template.ParseGlob("internal\\web\\*.gohtml") //получение всех gohtml файлов
	if err != nil {
		log.Fatalln(err)
	}
	return &Router{Router: mux.NewRouter(), services: services, tpl: tpl}
}

func (r *Router) Configure_router() { //Настройка роутера
	r.Router.HandleFunc("/ping", r.ping())
}

func (router *Router) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}
