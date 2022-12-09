package chat_database

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/config"
	"chat-app/internal/infrastructure/database"
	"chat-app/internal/user"
	"chat-app/internal/user/user_database"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	cfg      = config.New_config()
	db, err  = database.NewDB(*cfg)
	r        = user_database.NewUserRepoImpl(db)
	TestUser = user.NewUser("test1", "test1", "testemail4")
	chatR    = NewChatRepoImpl(db)
)

func truncTable(table string) {
	if _, err := db.Exec(fmt.Sprintf(fmt.Sprintf("TRUNCATE %s CASCADE", table))); err != nil {
		log.Fatal(err)
	}
}

func TestCreateChat(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, err := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	assert.NoError(t, err)
	assert.NotNil(t, chat)
	truncTable("users")
	truncTable("chats")
}

func TestGetChat(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))
	getChat, err := chatR.GetChat(chat.ID)

	assert.NoError(t, err)
	assert.NotNil(t, getChat)
	truncTable("users")
	truncTable("chats")

}

func TestGetChats(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	GetChats, err := chatR.GetChats(nil)

	assert.NoError(t, err)
	assert.NotNil(t, GetChats)
	truncTable("users")
	truncTable("chats")
}

func TestGetChatsByIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))
	chat1, _ := chatR.CreateChat(*chat_domain.NewChat("testName1", "testDescription1", u.ID, time.Now()))
	chat2, _ := chatR.CreateChat(*chat_domain.NewChat("testName2", "testDescription2", u.ID, time.Now()))

	filter := chat_domain.ChatFilter{IDs: []uint64{chat.ID, chat1.ID, chat2.ID}}

	GetChats, err := chatR.GetChats(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetChats)
	truncTable("users")
	truncTable("chats")
}

func TestGetChatsBySearch(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	desc := "testDescription"

	filter := chat_domain.ChatFilter{Search: &desc}

	GetChats, err := chatR.GetChats(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetChats)
	truncTable("users")
	truncTable("chats")
}

func TestGetChatsByUserIDs(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	filter := chat_domain.ChatFilter{UserIDs: []uint64{u.ID}}

	GetChats, err := chatR.GetChats(&filter)

	assert.NoError(t, err)
	assert.NotNil(t, GetChats)
	truncTable("users")
	truncTable("chats")
}

func TestUpdateChat(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, err := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	getChat, err := chatR.GetChat(chat.ID)

	getChat.Name = "111111"

	getChat, err = chatR.UpdateChat(*getChat)

	assert.NoError(t, err)
	assert.NotNil(t, getChat)
	truncTable("users")
	truncTable("chats")
}

func TestDeleteChat(t *testing.T) {
	u, _ := r.CreateUser(*TestUser)
	chat, _ := chatR.CreateChat(*chat_domain.NewChat("testName", "testDescription", u.ID, time.Now()))

	err := chatR.DeleteChat(chat.ID)

	assert.NoError(t, err)

	truncTable("users")

}
