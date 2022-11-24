package store_test

import (
	"chat/models"
	"chat/pkg/store"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_userrepository_Create(t *testing.T) {
	s, teardown := store.Test_store(t, db_URL)
	defer teardown("users")

	u, err := s.User().Create(models.Test_user(t), models.TestUserCredential(t))

	if err != nil {
		logrus.Error(err.Error())
	}

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func Test_userrepository_Find(t *testing.T) {
	s, teardown := store.Test_store(t, db_URL)
	defer teardown("users")

	username := "test@email.com"
	_, err := s.User().Find_by_email(username)
	assert.Error(t, err)

	u := models.Test_user(t)
	uc := models.TestUserCredential(t)
	u.FirstName = username

	s.User().Create(u, uc)

	u, err = s.User().Find_by_email(username)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
