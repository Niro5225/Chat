package connector_http

import (
	"chat-app/internal/api/handlers/handler_error"
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/connection/connection_domain"
	"chat-app/internal/connector/connector_domain"
	"chat-app/internal/user/user_domain"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ConnectorHandlers struct {
	UserService      *user_domain.UserServiceImp
	ChatService      *chat_domain.ChatServiceImp
	MessageService   *chat_domain.MessageServiceImp
	ConnectorService connector_domain.Connector
	httpError        *handler_error.HttpError
}

func NewChatHandler(
	userService *user_domain.UserServiceImp,
	chatService *chat_domain.ChatServiceImp,
	messageService *chat_domain.MessageServiceImp,
	ConnectorService connector_domain.Connector,
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

func (ch *ConnectorHandlers) AddToRoom(connector *connector_domain.ConnectorImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
			return
		}
		ws.Close()

		userId := c.Query("userId")
		chatId := c.Query("chatId")
		uintUserId, err := strconv.ParseUint(userId, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		uintChatId, err := strconv.ParseUint(chatId, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(uintUserId)
		fmt.Println(uintChatId)

		user, err := ch.UserService.GetUser(uint64(uintUserId))
		if err != nil {
			ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
		}

		conn := connection_domain.NewConnectionImpl(ws, *user, &uintChatId)

		connector.AddConnection(uintUserId, uintChatId, conn)

		fmt.Printf("connector.Rooms: %v\n", connector.Rooms)
		fmt.Printf("connector.Connections: %v\n", connector.Connections)

		// for {
		// 	//Read Message from client
		// 	mt, message, err := ws.ReadMessage()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// 	//If client message is ping will return pong
		// 	if string(message) == "ping" {
		// 		message = []byte("pong")
		// 	}
		// 	//Response message to client
		// 	err = ws.WriteMessage(mt, message)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// }

		// connector := connector_domain.NewConnectorImpl()

		// if err := connector.AddConnection(connection); err != nil {
		// 	ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"message":    "OK",
		// 	"connection": connection,
		// })
	}
}
