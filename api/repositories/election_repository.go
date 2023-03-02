package repositories

import (
	"chain-vote-api/models"
	"fmt"
)

func SaveElection(election *models.Election) (*models.Election, error) {
	err := DB.Create(&election).Error

	if err != nil {
		fmt.Println("Error saving election ")
		return nil, err
	}

	return election, nil
}

func GetElectionById(id string) (*models.Election, error) {
	election := models.Election{}

	err := DB.Where("id = ?", id).Preload("CreatedBy").First(&election).Error

	if err != nil {
		fmt.Println("Error getting election ")
		return nil, err
	}

	return &election, nil
}

func GetAllElections() ([]models.Election, error) {
	elections := []models.Election{}

	err := DB.Preload("CreatedBy").Find(&elections).Error

	if err != nil {
		fmt.Println("Error getting elections ")
		return nil, err
	}

	return elections, nil
}