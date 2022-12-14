package userhttp

import (
	"chat-app/internal/user/user_domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandlers) UserIdentity() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			uh.httpError.NewError(c, http.StatusUnauthorized, "Empty header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			uh.httpError.NewError(c, http.StatusUnauthorized, "Invalid handler")
			return
		}

		userId, err := user_domain.ParsToken(headerParts[1])
		if err != nil {
			uh.httpError.NewError(c, http.StatusUnauthorized, err.Error())
			return
		}

		uses, err := uh.UserService.GetUser(userId)
		if err != nil {
			uh.httpError.NewError(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("user", uses)
		c.Next()
	}
}
