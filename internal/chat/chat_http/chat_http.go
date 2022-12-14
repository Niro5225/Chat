package chathttp

import (
	"chat-app/internal/api/handlers/handler_error"
	"chat-app/internal/chat/chat_domain"
	chatdto "chat-app/internal/chat/chat_dto"
	"chat-app/internal/user/user_domain"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func fromDto(chatDto *chatdto.ChatDTO) chat_domain.Chat {
	return chat_domain.Chat{ID: chatDto.Id, Name: chatDto.Name, Description: chatDto.Description}
}

func toDto(chat *chat_domain.Chat) chatdto.ChatDTO {
	return chatdto.ChatDTO{Id: chat.ID, Name: chat.Name, Description: chat.Description}
}

type ChatHandlers struct {
	UserService    *user_domain.UserServiceImp
	ChatService    *chat_domain.ChatServiceImp
	MessageService *chat_domain.MessageServiceImp
	httpError      *handler_error.HttpError
}

func NewChatHandler(userService *user_domain.UserServiceImp, chatService *chat_domain.ChatServiceImp, messageService *chat_domain.MessageServiceImp, httpError *handler_error.HttpError) *ChatHandlers {
	return &ChatHandlers{UserService: userService, ChatService: chatService, MessageService: messageService, httpError: httpError}
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
			ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
		}
		chatDtos := []chatdto.ChatDTO{}

		for _, chat := range chats {
			chatDtos = append(chatDtos, toDto(&chat))
		}

		c.JSON(http.StatusOK, gin.H{
			"chats": chatDtos,
		})
		return
	}

	chats, err := ch.ChatService.GetChats(&filter)
	if err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	chatDtos := []chatdto.ChatDTO{}

	for _, chat := range chats {
		chatDtos = append(chatDtos, toDto(&chat))
	}

	c.JSON(http.StatusOK, gin.H{
		"chats": chatDtos,
	})

}

func (ch *ChatHandlers) ChatsId(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	chat, err := ch.ChatService.GetChat(uint64(uintId))
	if err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	chatDto := toDto(chat)

	c.JSON(http.StatusOK, gin.H{
		"message": chatDto,
	})
}

func (ch *ChatHandlers) CreateChat(c *gin.Context) {
	id, _ := c.Get("userId")
	fmt.Println(id)
	u, _ := ch.UserService.CreateUser(*user_domain.NewUser("testFirstName", "TestLastName", "testHandlerEmail"))
	chat := chat_domain.NewChat("testName", "testDescription", u.ID, time.Now())
	chat, err := ch.ChatService.CreateChat(*chat)
	if err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	chatDto := toDto(chat)

	c.JSON(http.StatusOK, gin.H{
		"message": chatDto,
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
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	chat.Name = "ChangedName"

	chat, err = ch.ChatService.UpdateChat(*chat)
	if err != nil {
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}

	chatDto := toDto(chat)

	c.JSON(http.StatusOK, gin.H{
		"message": chatDto,
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
		ch.httpError.NewError(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
