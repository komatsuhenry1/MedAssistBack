package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserId(c *gin.Context) string {
	claims, exists := c.Get("claims")
	if !exists {
		return ""
	}
	userId, ok := claims.(jwt.MapClaims)["sub"].(string)
	if !ok {
		return ""
	}
	return userId
}
