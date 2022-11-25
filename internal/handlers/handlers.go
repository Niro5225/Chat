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
}

func (r *Router) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}
