package middleware

import (
	"go-project/internal/jwt"
	"go-project/pkg/apperror"
	"go-project/pkg/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
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

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// AuthenticationMiddleware()
		userId, err := c.Get("userID")

		if !err {
			response.Response(c, nil, apperror.ErrUnauthorized.WithMessage("The token is not valid"))
			c.Abort()
			return
		}

		id, ok := userId.(uint)
		if !ok {
			response.Response(c, nil, apperror.ErrUnauthorized.WithMessage("The token is not valid"))
			c.Abort()
			return
		}
		if strconv.Itoa(int(id)) != c.Param("id") {
			response.Response(c, nil, apperror.ErrUnauthorized.WithMessage("This user does not have authorization right !"))
			c.Abort()
			return
		}

		c.Next()
	}
}
