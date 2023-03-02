package services

import (
	"chain-vote-api/mappers"
	"chain-vote-api/models"
	"chain-vote-api/repositories"
	"chain-vote-api/security"
	"chain-vote-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := repositories.SaveUser(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func CurrentUser(c *gin.Context) {

	userId, err := security.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := repositories.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userListing := mappers.UserToListing(u)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userListing})
}

// RegisterETHAddress updates the user's eth address for the logged user
func RegisterETHAddress(c *gin.Context) {
	user, err := utils.GetUserFromRequestContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateInput := models.UpdateAddressInput{}
	err = c.ShouldBindJSON(&updateInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValidAddress := utils.ValidateEthAddress(updateInput.EthAddress)

	if !isValidAddress {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid eth address"})
		return
	}

	err = repositories.UpdateUser(user, updateInput.EthAddress)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
