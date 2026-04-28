package middleware

import (
	"go-project/internal/jwt"
	"go-project/pkg/apperror"
	"go-project/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Response(c, nil, apperror.ErrUnauthorized.WithMessage("The token is empty"))
			c.Abort()
			return
		}

		user, err := jwt.Parse(token)

		if err != nil {
			response.Response(c, nil, apperror.ErrUnauthorized.WithMessage("The token is not valid"))
			c.Abort()
			return
		}
		log.Print(user.UserID)
		c.Set("userID", user.UserID)

		c.Next()
	}
}
