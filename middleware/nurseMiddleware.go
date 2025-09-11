
package middleware

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthNurse() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token not found",
				"success": false,
			})
			return
		}

		tokenString := header[len(BearerSchema):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signature method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
				"success": false,
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
				"success": false,
			})
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != "NURSE" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "restricted access to nurse",
				"success": false,
			})
			return
		}

		hidden, ok := claims["hidden"].(bool)
		if !ok || hidden {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "restricted access to hidden users",
				"success": false,
			})
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}
