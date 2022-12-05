package userhttp

import (
	"chat-app/internal/chat/chat_domain"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	UserService    *chat_domain.UserServiceImp
	ChatService    *chat_domain.ChatServiceImp
	MessageService *chat_domain.MessageServiceImp
}

func NewUserHandlers(userService *chat_domain.UserServiceImp, chatService *chat_domain.ChatServiceImp, messageService *chat_domain.MessageServiceImp) *UserHandlers {
	return &UserHandlers{UserService: userService, ChatService: chatService, MessageService: messageService}
}

func (uh *UserHandlers) GetUserId(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	u, err := uh.UserService.GetUser(uint64(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": u,
	})
}

func (uh *UserHandlers) GetUsers(c *gin.Context) {
	filter := chat_domain.UserFilter{}
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
	} else if query["email"] != nil {
		filter.Email = &query["email"][0]
	} else if query["search"] != nil {
		filter.Search = &query["search"][0]
	} else {
		users, err := uh.UserService.GetUsers(nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		for _, user := range users {
			fmt.Println(user)
		}
		return
	}

	users, err := uh.UserService.GetUsers(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func (uh *UserHandlers) GetMessages(c *gin.Context) {
	filter := chat_domain.MessageFilter{}
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
	} else if query["chatids"] != nil {
		Ids := strings.Split(query["chatids"][0], ",")
		for _, id := range Ids {
			uintId, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			filter.IDs = append(filter.ChatIDs, uint64(uintId))
		}
	} else if query["userids"] != nil {
		Ids := strings.Split(query["userids"][0], ",")
		for _, id := range Ids {
			uintId, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			filter.IDs = append(filter.UserIDs, uint64(uintId))
		}
	}
	messages, err := uh.MessageService.GetMessages(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
