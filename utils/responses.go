package utils

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
	})
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param : %s (type: %s) is required", name, typ)
}
