package mappers

import "chain-vote-api/models"

func ElectionToListing(election *models.Election) *models.ElectionListing {
	return &models.ElectionListing{
		ID:          election.ID,
		Name:        election.Name,
		Description: election.Description,
		StartDate:   election.StartDate,
		EndDate:     election.EndDate,
		Status:      election.Status.String(),
		Creator:     *UserToListing(&election.CreatedBy),
	}
}

func UserToListing(user *models.User) *models.UserListing {
	return &models.UserListing{
		ID:         user.ID,
		Username:   user.Username,
		ETHAddress: user.ETHAddress,
	}
}
