package services

import (
	"fmt"
	"gallery/models"
)

func CreateReaction(accountId uint, newReaction *models.Reaction) (reaction *models.Reaction, err error) {
	reaction = newReaction
	reaction.AccountId = accountId

	err = DB.Create(&reaction).Error
	if err != nil {
		return
	}
	return
}

func DeleteReaction(accountId uint, reactionId uint) (err error) {
	reaction := &models.Reaction{}
	err = DB.Where("photo_id = ? and account_id = ?", reactionId, accountId).First(reaction).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	err = DB.Delete(reaction).Error

	return
}
