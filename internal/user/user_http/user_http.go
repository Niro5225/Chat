package userhttp

import (
	"chat-app/internal/chat/chat_domain"
	"chat-app/internal/user/user_domain"
	userdto "chat-app/internal/user/user_dto"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserErrorResponse struct {
	Message string `json:"message"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, UserErrorResponse{message})
}

func fromDto(userDto *userdto.UserDTO) user_domain.User {
	return user_domain.User{ID: userDto.Id, FirstName: userDto.FirstName, LastName: userDto.LastName, Email: userDto.Email}
}

func toDto(user *user_domain.User) userdto.UserDTO {
	return userdto.UserDTO{Id: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}
}

type UserHandlers struct {
	UserService    *user_domain.UserServiceImp
	ChatService    *chat_domain.ChatServiceImp
	MessageService *chat_domain.MessageServiceImp
}

func NewUserHandlers(userService *user_domain.UserServiceImp, chatService *chat_domain.ChatServiceImp, messageService *chat_domain.MessageServiceImp) *UserHandlers {
	return &UserHandlers{UserService: userService, ChatService: chatService, MessageService: messageService}
}

func (uh *UserHandlers) GetUserId(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Fatal(err)
		return
	}
	u, err := uh.UserService.GetUser(uint64(uintId))
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())

		return
	}

	userDto := toDto(u)

	c.JSON(http.StatusOK, gin.H{
		"message": userDto,
	})
}

func (uh *UserHandlers) GetUsers(c *gin.Context) {
	filter := user_domain.UserFilter{}
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
			NewError(c, http.StatusBadRequest, err.Error())
		}
		userDtos := []userdto.UserDTO{}

		for _, user := range users {
			userDtos = append(userDtos, toDto(&user))
		}
		c.JSON(http.StatusOK, gin.H{
			"users": userDtos,
		})
		return
	}

	users, err := uh.UserService.GetUsers(&filter)
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())

	}

	userDtos := []userdto.UserDTO{}

	for _, user := range users {
		userDtos = append(userDtos, toDto(&user))
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userDtos,
	})

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
	} else {
		messages, err := uh.MessageService.GetMessages(nil)
		if err != nil {
			NewError(c, http.StatusBadRequest, err.Error())

		}
		for _, message := range messages {
			fmt.Println(message)
		}
		return
	}
	messages, err := uh.MessageService.GetMessages(&filter)
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())

	}
	for _, message := range messages {
		fmt.Println(message)
	}
}

func (uh *UserHandlers) Login(c *gin.Context) {
	inputData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&inputData); err != nil {
		NewError(c, http.StatusBadRequest, err.Error())

		return
	}

	user, token, err := uh.UserService.SignIn(inputData.Email, inputData.Password)
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	userDto := toDto(user)

	c.JSON(http.StatusOK, gin.H{
		"user":  userDto,
		"token": token,
	})

}

func (uh *UserHandlers) Registration(c *gin.Context) {
	var input user_domain.User

	if err := c.BindJSON(&input); err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := uh.UserService.CreateUser(*user_domain.NewUser(input.FirstName, input.LastName, input.Email))
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	userDto := toDto(user)

	uc, err := uh.UserService.CreateUserCredential(*user_domain.NewUserCredential(user.ID, "testRegPassword", user.Email))
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	user, token, err := uh.UserService.SignUp(*user, *uc)
	if err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  userDto,
		"token": token,
	})

}
