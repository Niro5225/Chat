package connection_domain

import (
	"chat-app/internal/user/user_domain"

	"github.com/gorilla/websocket"
)

type Connection interface {
	SendMessage(data interface{}) error
}

type ConnectionImpl struct {
	wsConn        *websocket.Conn
	userID        user_domain.User
	currentChatID *uint64
}

func NewConnectionImpl(wsConn *websocket.Conn, userID user_domain.User, currentChatID *uint64) *ConnectionImpl {
	return &ConnectionImpl{wsConn: wsConn, userID: userID, currentChatID: currentChatID}
}

func (c *ConnectionImpl) SendMessage(data interface{}) error {
	return nil
}
