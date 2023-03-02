package services

import (
	"chain-vote-api/enums"
	"chain-vote-api/mappers"
	"chain-vote-api/models"
	"chain-vote-api/repositories"
	"chain-vote-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateElection (c *gin.Context) {
	input := models.CreateElectionInput{}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creatorId,err := utils.GetUserIdFromRequestContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := models.Election{}
	e.Name = input.Name
	e.Description = input.Description
	e.StartDate = input.StartDate
	e.EndDate = input.EndDate
	e.Status = enums.PENDING
	e.CreatedByID = creatorId
	
	saved,err := repositories.SaveElection(&e)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	electionListing := mappers.ElectionToListing(saved)

	c.JSON(http.StatusOK, gin.H{"data": electionListing})
}

func GetElectionById(c *gin.Context) {
	id := c.Param("id")
	e,err := repositories.GetElectionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	eListing := mappers.ElectionToListing(e)
	
	c.JSON(http.StatusOK, gin.H{"data": eListing})
}

func GetAllElections(c *gin.Context) {
	elections,err := repositories.GetAllElections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	electionsListing := make([]*models.ElectionListing,0)

	for _,e := range elections {
		electionsListing = append(electionsListing, mappers.ElectionToListing(&e))
	}

	c.JSON(http.StatusOK, gin.H{"data": electionsListing})
}