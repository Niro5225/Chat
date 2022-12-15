package connector_database

import "chat-app/internal/connection/connection_domain"

type ConnectorImpl struct {
	Rooms       map[uint64][]connection_domain.Connection // chatID -> []Connection
	Connections map[uint64][]connection_domain.Connection // userID -> []Connection

}

func NewConnectorImpl(Rooms map[uint64][]connection_domain.Connection, Connections map[uint64][]connection_domain.Connection) *ConnectorImpl {
	return &ConnectorImpl{Rooms: Rooms, Connections: Connections}
}

func (r *ConnectorImpl) AddConnection(conn connection_domain.Connection) error {
	return nil
}

func (r *ConnectorImpl) SendMessageByRoom(roomID uint64, msg interface{}) {
	return
}

func (r *ConnectorImpl) GetUserConnection(userID uint64) *connection_domain.Connection {
	return nil
}
