package api

import (
	"fmt"
	"owwi/pkg/models"
	"owwi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Create Partner
// @Description Create Partner
// @Tags Partner
// @Accept json
// @Produce json
// @Param register body models.CreatePartner true "Partner data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /partners [post]
func createPartner(c *gin.Context) {
	var partnerData models.CreatePartner
	if err := c.BindJSON(&partnerData); err != nil || partnerData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	partnerData.User = userID.(string)

	if err := repositories.PartnerRepository.CreatePartner(partnerData); err != nil {
		fmt.Print(err.Error())
		c.JSON(500, gin.H{"error": "Failed to create type"})
		return
	}

	c.Status(201)
}

// @Summary Update Partner
// @Description Update Partner By ID
// @Tags Partner
// @Accept json
// @Produce json
// @Param register body models.UpdatePartner true "Partner data"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /partners [put]
func updatePartner(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Partner ID is required"})
		return
	}
	var partnerData models.UpdatePartner
	if err := c.BindJSON(&partnerData); err != nil || partnerData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	partnerData.ID = id
	if err := repositories.PartnerRepository.UpdatePartner(partnerData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update partner"})
		return
	}

	c.Status(200)
}

// @Summary Get Partner By ID
// @Description Get Partner By ID
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner ID"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /partners/:id [get]
func getPartnerByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Partner ID is required"})
		return
	}

	partner, err := repositories.PartnerRepository.GetPartnerByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve partner", "err": err.Error()})
		return
	}

	if partner == nil {
		c.JSON(404, gin.H{"error": "Partner not found"})
		return
	}

	c.JSON(200, partner)
}

// @Summary Get All Partners By User
// @Description Get all partners associated with the authenticated user
// @Tags Partner
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /partners [get]
func getAllPartnersByUser(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	partners, err := repositories.PartnerRepository.GetAllPartnersByUser(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve partners"})
		return
	}

	if len(partners) == 0 {
		c.JSON(404, gin.H{"message": "No partners found for this user", "user_id": userID, "partners": partners})
		return
	}

	c.JSON(200, partners)
}
// @Summary Delete Partner
// @Description Delete Partner By ID
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner ID"
// @Security BearerAuth
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /partners/:id [delete]
func deletePartner(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Partner ID is required"})
		return
	}

	if err := repositories.PartnerRepository.DeletePartner(id); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete partner"})
		return
	}

	c.Status(204)
}

var PartnerApi = struct {
	CreatePartner         func(*gin.Context)
	UpdatePartner         func(*gin.Context)
	GetPartnerByID        func(*gin.Context)
	GetAllPartnersByUser func(*gin.Context)
	DeletePartner         func(*gin.Context)
}{
	CreatePartner:         createPartner,
	UpdatePartner:         updatePartner,
	GetPartnerByID:        getPartnerByID,
	GetAllPartnersByUser: getAllPartnersByUser,
	DeletePartner:         deletePartner,
}
