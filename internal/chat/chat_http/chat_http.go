package chathttp

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user/user_domain"

	"chat-app/internal/user"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatHandlers struct {
	UserService    *user_domain.UserServiceImp
	ChatService    *chat_domain.ChatServiceImp
	MessageService *chat_domain.MessageServiceImp
}

func NewChatHandler(userService *user_domain.UserServiceImp, chatService *chat_domain.ChatServiceImp, messageService *chat_domain.MessageServiceImp) *ChatHandlers {
	return &ChatHandlers{UserService: userService, ChatService: chatService, MessageService: messageService}
}

func (ch *ChatHandlers) GetChatsQuery(c *gin.Context) {
	filter := chat_domain.ChatFilter{}
	query := c.Request.URL.Query()

	if query["ids"] != nil {
		Ids := strings.Split(query["ids"][0], ",")
		for _, id := range Ids {
			uintId, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			filter.IDs = append(filter.IDs, uint64(uintId))
		}
	} else if query["search"] != nil {
		filter.Search = &query["search"][0]
	} else if query["userids"] != nil {
		Ids := strings.Split(query["userids"][0], ",")
		for _, id := range Ids {
			uintId, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			filter.IDs = append(filter.UserIDs, uint64(uintId))
		}
	} else {
		chats, err := ch.ChatService.GetChats(nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		for _, chat := range chats {
			fmt.Println(chat)
		}
		return
	}

	chats, err := ch.ChatService.GetChats(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	for _, chat := range chats {
		fmt.Println(chat)
	}

}

func (ch *ChatHandlers) ChatsId(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	chat, err := ch.ChatService.GetChat(uint64(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": chat,
	})
}

func (ch *ChatHandlers) CreateChat(c *gin.Context) {
	u, _ := ch.UserService.CreateUser(*user.NewUser("testFirstName", "TestLastName", "testHandlerEmail"))
	chat := chat_domain.NewChat("testName", "testDescription", u.ID, time.Now())
	chat, err := ch.ChatService.CreateChat(*chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": chat,
	})
}

func (ch *ChatHandlers) UpdateChat(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	chat, err := ch.ChatService.GetChat(uint64(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	chat.Name = "ChangedName"

	chat, err = ch.ChatService.UpdateChat(*chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": chat,
	})
}

func (ch *ChatHandlers) DeleteChat(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.ChatService.DeleteChat(uint64(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
