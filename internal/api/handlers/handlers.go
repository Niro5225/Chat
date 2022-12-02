package handlers

import (
	"chat-app/internal/chat/chat_domain"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router         *mux.Router
	UserService    *chat_domain.UserServiceImp
	ChatService    *chat_domain.ChatServiceImp
	MessageService *chat_domain.MessageServiceImp
	tpl            *template.Template
}

func NewRouter(userService *chat_domain.UserServiceImp, chatService *chat_domain.ChatServiceImp, messageService *chat_domain.MessageServiceImp) *Router { //Создание роутера
	var tpl *template.Template

	tpl, err := template.ParseGlob("internal\\web\\*.gohtml") //получение всех gohtml файлов
	if err != nil {
		log.Fatalln(err)
	}
	return &Router{Router: mux.NewRouter(), UserService: userService, ChatService: chatService, MessageService: messageService, tpl: tpl}
}

func (r *Router) Configure_router() { //Настройка роутера
	r.Router.HandleFunc("/ping", r.ping())
	r.Router.HandleFunc("/", r.test())
	r.Router.HandleFunc("/t", r.test1())

}

func (router *Router) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func (router *Router) test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := router.UserService.CreateUser(*chat_domain.NewUser("testFirstName", "testLastName", "testEmail"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		uc, err := router.UserService.CreateUserCredential(*chat_domain.NewUserCredential(u.ID, "testPssword", u.Email))
		if err != nil {
			io.WriteString(w, err.Error())
		}

		su, token, err := router.UserService.SignUp(*u, *uc)

		data := struct {
			fname string
			lname string
			email string
			token string
		}{su.FirstName, su.LastName, su.Email, token}

		router.tpl.ExecuteTemplate(w, "index.gohtml", data)
	}
}

func (router *Router) test1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := "testEmail"
		si, token, err := router.UserService.SignIn(email, "testPssword1")
		if err != nil {
			io.WriteString(w, err.Error())
		}
		data := struct {
			fname string
			lname string
			email string
			token string
		}{si.FirstName, si.LastName, si.Email, token}

		router.tpl.ExecuteTemplate(w, "index.gohtml", data)
	}
}
