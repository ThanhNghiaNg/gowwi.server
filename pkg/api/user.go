package api

import (
	"github.com/gin-gonic/gin"
	"owwi/pkg/models" // Import your models package
)


func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(201, user)
}