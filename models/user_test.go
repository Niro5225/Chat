package models_test

import (
	"chat/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_user_before_create(t *testing.T) {
	u := models.Test_user(t)
	assert.NoError(t, u.Before_create())
	assert.NotEmpty(t, u.Encrypted_Password)
}

func Test_validate(t *testing.T) {
	u := models.Test_user(t)
	assert.NoError(t, u.Validate())
}
