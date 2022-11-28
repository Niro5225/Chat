package repository

import (
	"chat-app/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	messageR = NewMessageR(db)
)

func TestCreateMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, err := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))

	assert.NoError(t, err)
	assert.NotNil(t, message)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))

	getMessage, err := messageR.GetMessage(message.ID)

	assert.NoError(t, err)
	assert.NotNil(t, getMessage)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*models.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*models.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := models.MessageFilter{IDs: []uint64{message.ID, message1.ID, message2.ID}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesBySearch(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))

	filter := models.MessageFilter{Search: &message.Text}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByChatIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*models.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*models.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := models.MessageFilter{ChatIDs: []uint64{message.ChatID, message1.ChatID, message2.ChatID}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByUserIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*models.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*models.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := models.MessageFilter{UserIDs: []uint64{message.CreatedBy, message1.CreatedBy, message2.CreatedBy}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestUpdateMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))

	message.Text = "new message"

	updateMessage, err := messageR.UpdateMessage(*message)

	assert.NoError(t, err)
	assert.NotNil(t, updateMessage)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestDeleteMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*models.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*models.NewMessage("test text", chat.ID, u.ID, time.Now()))

	err := messageR.DeleteMessage(message.ID)

	assert.NoError(t, err)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}
