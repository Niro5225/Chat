package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	messageR = NewMessageRepoImpl(db)
)

func TestCreateMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, err := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))

	assert.NoError(t, err)
	assert.NotNil(t, message)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))

	getMessage, err := messageR.GetMessage(message.ID)

	assert.NoError(t, err)
	assert.NotNil(t, getMessage)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := chat_domain.MessageFilter{IDs: []uint64{message.ID, message1.ID, message2.ID}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesBySearch(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))

	filter := chat_domain.MessageFilter{Search: &message.Text}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByChatIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := chat_domain.MessageFilter{ChatIDs: []uint64{message.ChatID, message1.ChatID, message2.ChatID}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestGetMessagesByUserIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))
	message1, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text1", chat.ID, u.ID, time.Now()))
	message2, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text2", chat.ID, u.ID, time.Now()))

	filter := chat_domain.MessageFilter{UserIDs: []uint64{message.CreatedBy, message1.CreatedBy, message2.CreatedBy}}

	GetMessages, err := messageR.GetMessages(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetMessages)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}

func TestUpdateMessage(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))

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
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	message, _ := messageR.CreateMessage(*chat_domain.NewMessage("test text", chat.ID, u.ID, time.Now()))

	err := messageR.DeleteMessage(message.ID)

	assert.NoError(t, err)

	truncTable("users")
	truncTable("chats")
	truncTable("messages")
}
