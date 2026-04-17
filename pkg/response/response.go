package response

import (
	"errors"
	"go-project/pkg/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

// pkg/response/response.go

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"message":    "success",
		"data":       data,
	})
}

func Error(c *gin.Context, err error) {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, gin.H{
			"error_code": appErr.Code,
			"message":    appErr.Message,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}

func Response(c *gin.Context, data any, err error) {
	if err != nil {
		Error(c, err)
		return
	}
	Success(c, data)
}
