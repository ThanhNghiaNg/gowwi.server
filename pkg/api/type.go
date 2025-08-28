package api

import (
	"owwi/pkg/models"
	"owwi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Create Type
// @Description Create a new type
// @Tags Type
// @Accept json
// @Produce json
// @Param type body models.Type true "Type data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 500
// @Router /types [post]
func createType(c *gin.Context) {
	var typeData models.Type;
	if err := c.BindJSON(&typeData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}

	if err := repositories.TypeRepository.CreateType(typeData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create type"})
	}

	c.JSON(201, typeData)
}


func getTypes(c *gin.Context) {
	userID := c.GetString("user_id")
	types, err := repositories.TypeRepository.GetTypesByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve types"})
		return
	}

	c.JSON(200, types)
}

var TypeApi = struct {
	CreateType         func(*gin.Context)
	GetTypes          func(*gin.Context)
}{
	CreateType:         createType,
	GetTypes:          getTypes,
}
