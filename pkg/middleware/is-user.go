package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsUser(c *gin.Context) {
	var tokenString = c.Request.Header.Get("Authorization")

	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.JSON(401, gin.H{"error": "Parse token failed", "err": err.Error()})
		c.Abort()
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["user_id"] == "" || claims["username"] == "" || claims["role"] == "" {
			c.JSON(401, gin.H{"error": "Claims are not valid"})
			c.Abort()
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()
	} else {
		c.JSON(401, gin.H{"error": "Claims are not valid"})
		c.Abort()
	}
}
