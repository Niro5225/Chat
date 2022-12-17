package main

import (
	"chat-app/internal/api/handlers"
	"chat-app/internal/api/handlers/handler_error"
	"chat-app/internal/api/server"
	"chat-app/internal/chat/chat_database"
	"chat-app/internal/chat/chat_domain"
	chathttp "chat-app/internal/chat/chat_http"
	"chat-app/internal/config"
	"chat-app/internal/connector/connector_domain"
	"chat-app/internal/connector/connector_http"
	"chat-app/internal/infrastructure/database"
	"chat-app/internal/user/user_database"
	"chat-app/internal/user/user_domain"
	userhttp "chat-app/internal/user/user_http"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New_config() //обьект конфига

	db, err := database.NewDB(*cfg) //обьект бд
	if err != nil {
		logrus.Fatal(err)
	}

	HttpError := handler_error.NewHttpError()

	UserRepository := user_database.NewUserRepoImpl(db)
	ChatRepository := chat_database.NewChatRepoImpl(db)
	MessageRepository := chat_database.NewMessageRepoImpl(db)

	UserService := user_domain.NewUserServiceImp(UserRepository)
	CharService := chat_domain.NewChatServiceImp(ChatRepository)
	MessageService := chat_domain.NewMessageServiceImp(MessageRepository)
	ConnectorService := connector_domain.NewConnectorImpl()

	userHandlers := userhttp.NewUserHandlers(UserService, CharService, MessageService, HttpError)
	catHandlers := chathttp.NewChatHandler(UserService, CharService, MessageService, HttpError)
	connectorHandlers := connector_http.NewChatHandler(UserService, CharService, MessageService, ConnectorService, HttpError)

	router := handlers.NewRouter(userHandlers, catHandlers, connectorHandlers) //обьект роутера
	serv := server.NewApiServer(cfg, logrus.New(), router)                     //обьект сервера

	if err := serv.Start(); err != nil { //запуск сервера
		logrus.Fatal(err.Error())
	}

}
