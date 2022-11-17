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

func (a *Api_server) process() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		username := r.FormValue("Login")
		password := r.FormValue("Password")
		a.logger.Info(username, " ", password)

		//LOGIN
		// u, err := a.store.User().Find_by_username(username)
		// if err != nil {
		// 	a.logger.Error(err.Error())
		// }

		// a.logger.Info(u)

		// u, err = a.store.User().Login(u)
		// if err != nil {
		// 	a.logger.Error(err.Error())
		// 	return
		// }
		// io.WriteString(w, u.Token)

		//REG
		u, err := a.store.User().Create(models.New_ueser(username, password))
		if err != nil {
			a.logger.Error(err.Error())
		}

		a.logger.Info(u)

		return

	}
}

func (a *Api_server) get_id() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		username := r.FormValue("Login")
		password := r.FormValue("Password")

		u, err := a.store.User().Find_by_username(username)
		if err != nil {
			a.logger.Error(err.Error())
		}

		if c, _ := u.Check_password(password); !c {
			io.WriteString(w, "BAD PASSWORD")
		}

		io.WriteString(w, fmt.Sprintf("PASSWORD IS OK ID IS %d", u.Id))
	}
}
