package handlers

import (
	chathttp "chat-app/internal/chat/chat_http"
	"chat-app/internal/connector/connector_domain"
	"chat-app/internal/connector/connector_http"
	userhttp "chat-app/internal/user/user_http"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Router            *gin.Engine
	userHandler       *userhttp.UserHandlers
	chatHandlers      *chathttp.ChatHandlers
	connectorHandlers *connector_http.ConnectorHandlers
	connector         *connector_domain.ConnectorImpl
	tpl               *template.Template
}

func NewRouter(userHandlers *userhttp.UserHandlers,
	chatHandlers *chathttp.ChatHandlers,
	connectorHandlers *connector_http.ConnectorHandlers,
	connector *connector_domain.ConnectorImpl) *Router { //Создание роутера
	var tpl *template.Template

	tpl, err := template.ParseGlob("internal\\web\\*.gohtml") //получение всех gohtml файлов
	if err != nil {
		log.Fatalln(err)
	}
	return &Router{Router: gin.New(), tpl: tpl, chatHandlers: chatHandlers, userHandler: userHandlers, connectorHandlers: connectorHandlers, connector: connector}
}

func (r *Router) Configure_router() { //Настройка роутера
	r.Router.GET("/ping", r.ping)
	// r.Router.GET("/test", r.connectorHandlers.AddToRoom(r.connector))
	users := r.Router.Group("/users")
	users.Use(r.userHandler.UserIdentity())
	{
		users.GET("/", r.userHandler.GetUsers)
		users.GET("/:id", r.userHandler.GetUserId)
	}
	chats := r.Router.Group("/chats")
	chats.Use(r.userHandler.UserIdentity())
	{
		chats.GET("/", r.chatHandlers.GetChatsQuery)
		chats.GET("/:id", r.chatHandlers.ChatsId)
		chats.POST("/", r.chatHandlers.CreateChat)
		chats.PUT("/", r.chatHandlers.UpdateChat)
		chats.DELETE("/:id", r.chatHandlers.DeleteChat)
		chats.GET("/messages", r.userHandler.GetMessages)
	}

	ws := r.Router.Group("/ws")
	{
		ws.GET("/add_to_chat", r.connectorHandlers.AddToRoom(r.connector))
	}

	auth := r.Router.Group("/auth")
	{
		auth.POST("/signin", r.userHandler.Login)
		auth.POST("/signup", r.userHandler.Registration)
	}

}

func (router *Router) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}
