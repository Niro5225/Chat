package apiserver

import (
	"chat/models"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (a *Api_server) handle_hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}

func (a *Api_server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.store.User().Get_all_users()
		a.logger.Info(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := struct {
			Users  []models.User
			UserID int
		}{Users: users, UserID: 0}

		a.tpl.ExecuteTemplate(w, "index.gohtml", data)
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

func (a *Api_server) loginProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//LOGIN
		Email := r.FormValue("Email")
		password := r.FormValue("password")
		u, err := a.store.User().Find_by_email(Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		u, err = a.store.User().Login(u, password)
		if err != nil {
			// io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			users, err := a.store.User().Get_all_users()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			data := struct {
				Users  []models.User
				UserID int
			}{Users: users, UserID: int(u.ID)}
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			a.tpl.ExecuteTemplate(w, "index.gohtml", data)
		}

	}
}

func (a *Api_server) RegistrationProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		FirstName := r.FormValue("FirstName")
		LastName := r.FormValue("LastName")
		Email := r.FormValue("Email")
		Password := r.FormValue("password")

		u, err := a.store.User().Create(models.New_ueser(FirstName, LastName, Email), models.NewUserCredential(Password))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users, err := a.store.User().Get_all_users()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := struct {
			Users  []models.User
			UserID int
		}{Users: users, UserID: int(u.ID)}
		a.tpl.ExecuteTemplate(w, "index.gohtml", data)

	}
}

func (a *Api_server) CreateChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := r.FormValue("UserID")
		a.logger.Info(ID)
		data := struct {
			ID string
		}{ID: ID}

		a.tpl.ExecuteTemplate(w, "create_chat.gohtml", data)
	}
}

func (a *Api_server) CreateChatProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ChatName := r.FormValue("Chat_name")
		ChatDescription := r.FormValue("Description")
		UserId := r.FormValue("CreateChat")
		ID, _ := strconv.Atoi(UserId)
		ch, err := a.store.Chat().CreateChat(models.NewChat(ChatName, ChatDescription, ID, time.Now()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		a.logger.Info(ch)

		data := struct {
			Name     string
			Messages []models.Message
			ChatID   uint64
			UserID   int
		}{Name: ch.Name, Messages: []models.Message{}, ChatID: ch.ID, UserID: ID}
		a.tpl.ExecuteTemplate(w, "chat.gohtml", data)
	}
}

func (a *Api_server) FindEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := ""
		fmt.Scanln(&email)
		_, err := a.store.User().Find_by_email(email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a *Api_server) MessageProcess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messageText := r.FormValue("message")
		chatId := r.FormValue("CHATID")
		a.logger.Info(chatId)
		userId := r.FormValue("UserID")
		cID, _ := strconv.Atoi(chatId)
		uID, _ := strconv.Atoi(userId)
		m, err := a.store.Message().CreateMessage(models.NewMessage(messageText, uint64(cID), uint64(uID), time.Now()))
		if err != nil {
			a.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		a.logger.Info(m)

		a.logger.Info("IM HERE 1")

		ch, err := a.store.Chat().FindChatById(uint64(cID))
		if err != nil {
			a.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a.logger.Info("IM HERE 2")

		messages, err := a.store.Message().GetChatMessages(ch.ID)
		if err != nil {
			a.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		a.logger.Info("IM HERE 3")

		data := struct {
			Name     string
			Messages []models.Message
			ChatID   uint64
			UserID   int
		}{Name: ch.Name, Messages: messages, ChatID: ch.ID, UserID: uID}
		a.tpl.ExecuteTemplate(w, "chat.gohtml", data)
	}
}

func (a *Api_server) OpenChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type data struct {
		}
		a.tpl.ExecuteTemplate(w, "chat.gohtml", nil)
	}
}
