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
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Empty header",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid handler",
			})
			return
		}

		userId, err := user_domain.ParsToken(headerParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		uses, err := uh.UserService.GetUser(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("user", uses)
		c.Next()
	}
}
