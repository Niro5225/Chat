package connector_domain

import "chat-app/internal/connection/connection_domain"

type ConnectorRepository interface {
	AddConnection(conn connection_domain.Connection) error
	SendMessageByRoom(roomID uint64, msg interface{})
	GetUserConnection(userID uint64) *connection_domain.Connection
}
