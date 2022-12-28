package connector_domain

import "chat-app/internal/connection/connection_domain"

type Connector interface {
	AddConnection(UserId uint64, ChatId uint64, conn connection_domain.Connection)
	SendMessageByRoom(roomID uint64, msg interface{})
	GetUserConnection(userID uint64) *connection_domain.Connection
}

type ConnectorImpl struct {
	Rooms       map[uint64][]connection_domain.Connection // chatID -> []Connection
	Connections map[uint64][]connection_domain.Connection // userID -> []Connection
}

func NewConnectorImpl() *ConnectorImpl {
	return &ConnectorImpl{
		Rooms:       make(map[uint64][]connection_domain.Connection),
		Connections: make(map[uint64][]connection_domain.Connection),
	}
}

func (c *ConnectorImpl) AddConnection(UserId uint64, ChatId uint64, conn connection_domain.Connection) {
	// id := len(c.Rooms) + 1
	c.Rooms[ChatId] = []connection_domain.Connection{conn}
	c.Connections[UserId] = []connection_domain.Connection{conn}
}

func (c *ConnectorImpl) SendMessageByRoom(roomID uint64, msg interface{}) {

}

func (c *ConnectorImpl) GetUserConnection(userID uint64) *connection_domain.Connection {

	for id, con := range c.Connections {
		if userID == id {
			return &con[0]
		}
	}
	return nil
}
