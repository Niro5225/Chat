package connection_domain

type ConnectionRepository interface {
	SendMessage(data interface{}) error
}
