package connection_domain

type Connection interface {
	SendMessage(data interface{}) error
}

type ConnectionService struct {
	repo ConnectionRepository
}

func NewConnectionService(repo ConnectionRepository) *ConnectionService {
	return &ConnectionService{repo: repo}
}

func (s *ConnectionService) SendMessage(data interface{}) error {
	return s.repo.SendMessage(data)
}
