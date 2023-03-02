package utils

import (
	"chain-vote-api/models"
	"chain-vote-api/repositories"
	"chain-vote-api/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

// ValidateEthAddress Determine if the string is a valid ethereum blockchain address
func ValidateEthAddress(ethAddress string) bool {
	ethRegex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$") //https://goethereumbook.org/en/address-check/
	return ethRegex.MatchString(ethAddress)
}

// GetUserFromRequestContext gets user id from token and query's the database for the user, request must be authenticated
func GetUserFromRequestContext(context *gin.Context) (*models.User, error) {
	userId, err := GetUserIdFromRequestContext(context)

	if err != nil {
		return nil, err
	}

	user, err := repositories.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserIdFromRequestContext(context *gin.Context) (uuid.UUID, error) {
	userId, err := security.ExtractTokenID(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return uuid.Nil, err
	}
	return userId, nil
}
