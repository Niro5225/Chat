package store_test

import (
	"chat/models"
	"chat/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_messagerepository_Create(t *testing.T) {
	s, teardown := store.Test_store(t, db_URL)
	defer teardown("messages")

	// _, err := s.User().Create(models.Test_user(t), models.TestUserCredential(t))
	// if err != nil {
	// 	logrus.Error(err.Error())
	// }
	// ch, err := s.Chat().CreateChat(models.TestChat(t))

	// if err != nil {
	// 	logrus.Error(err.Error())
	// }

	m, err := s.Message().CreateMessage(models.TestMessage(t))

	assert.NoError(t, err)
	assert.NotNil(t, m)
}
