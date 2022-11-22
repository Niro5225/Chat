package apiserver

import (
	"chat/models"
	"fmt"
	"io"
	"net/http"
)

func (a *Api_server) handle_hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func (a *Api_server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}
}

func (a *Api_server) registration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.tpl.ExecuteTemplate(w, "registration.gohtml", nil)
	}
}

func (a *Api_server) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
}

func (a *Api_server) login_process() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//LOGIN
		Email := r.FormValue("Email")
		u, err := a.store.User().Find_by_email(Email)
		if err != nil {
			a.logger.Error(err.Error())
		}

		a.logger.Info(u)

		_, err = a.store.User().Login(u)
		if err != nil {
			io.WriteString(w, string(http.StatusBadRequest))
			a.logger.Error(err.Error())
			return
		}

	}
}

func (a *Api_server) registration_process() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		FirstName := r.FormValue("FirstName")
		LastName := r.FormValue("LastName")
		Email := r.FormValue("Email")
		//REG
		u, err := a.store.User().Create(models.New_ueser(FirstName, LastName, Email))
		if err != nil {
			a.logger.Error(err.Error())
		}

		a.logger.Info(u)

		return

	}
}

func (a *Api_server) find_email() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := ""
		fmt.Scanln(&email)
		_, err := a.store.User().Find_by_email(email)
		if err != nil {
			a.logger.Error(err)
		}
	}
}
