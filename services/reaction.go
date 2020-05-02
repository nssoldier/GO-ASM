package services

import (
	"errors"
	"fmt"
	"gallery/models"
)

func CreateReaction(accountId uint, newReaction *models.Reaction) (reaction *models.Reaction, err error) {
	reaction = newReaction
	reaction.AccountId = accountId
	photo := &models.Photo{}
	err = DB.Where("id = ?", reaction.PhotoId).First(photo).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = GetPublicGalleryById(photo.GalleryId)

	if err != nil {
		err = errors.New("can not find this public photo")
		return
	}
	err = DB.Create(&reaction).Error
	if err != nil {
		err = errors.New("can not find this public photo")
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
