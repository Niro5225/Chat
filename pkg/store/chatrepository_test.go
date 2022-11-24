package store_test

import (
	"chat/models"
	"chat/pkg/store"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_chatrepository_Create(t *testing.T) {
	s, teardown := store.Test_store(t, db_URL)
	defer teardown("chats")

	u, err := s.User().Create(models.Test_user(t), models.TestUserCredential(t))
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(u)

	ch, err := s.Chat().CreateChat(models.TestChat(t))

	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.Info(ch)

	assert.NoError(t, err)
	assert.NotNil(t, ch)
}
