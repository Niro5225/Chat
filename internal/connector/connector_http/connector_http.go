package connector_http

import (
	"chat-app/internal/api/handlers/handler_error"
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/connection/connection_domain"
	"chat-app/internal/connector/connector_domain"
	"chat-app/internal/user/user_domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ConnectorHandlers struct {
	UserService      *user_domain.UserServiceImp
	ChatService      *chat_domain.ChatServiceImp
	MessageService   *chat_domain.MessageServiceImp
	ConnectorService *connector_domain.ConnectorImpl
	httpError        *handler_error.HttpError
}

func NewChatHandler(
	userService *user_domain.UserServiceImp,
	chatService *chat_domain.ChatServiceImp,
	messageService *chat_domain.MessageServiceImp,
	ConnectorService *connector_domain.ConnectorImpl,
	httpError *handler_error.HttpError) *ConnectorHandlers {
	return &ConnectorHandlers{UserService: userService, ChatService: chatService, MessageService: messageService, ConnectorService: ConnectorService, httpError: httpError}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (ch *ConnectorHandlers) SendMessageByRoom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SendMessageByRoom",
	})
}

func (ch *ConnectorHandlers) AddToRoom(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "AddToRoom method",
	// })
	inputData := struct {
		UserId uint64 `json:"user_id"`
		ChatId uint64 `json:"chat_id"`
	}{}

	if err := c.BindJSON(&inputData); err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())

		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	user, err := ch.UserService.GetUser(inputData.UserId)
	if err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	con := connection_domain.NewConnectionImpl(ws, *user, &inputData.ChatId)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"con":     con,
	})
}
