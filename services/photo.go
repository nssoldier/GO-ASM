package services

import (
	"fmt"
	"gallery/models"
)

func CreatePhoto(accountId uint, newPhoto *models.Photo) (photo *models.Photo, err error) {
	photo = newPhoto
	photo.AccountId = accountId

	err = DB.Create(&photo).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetPhotoById(photoId uint) (photo *models.Photo, err error) {

	photo = &models.Photo{}
	err = DB.First(photo, photoId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	reactions, err := getReactionsByPhoto(*photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	photo.Reactions = reactions

	return
}

func GetPublicPhotoById(photoId uint) (photo *models.Photo, err error) {

	photo = &models.Photo{}
	err = DB.First(photo, photoId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = GetPublicGalleryById(photo.GalleryId)

	if err != nil {
		fmt.Println(err)
		return
	}

	reactions, err := getReactionsByPhoto(*photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	photo.Reactions = reactions

	return
}

func UpdatePhoto(accountId uint, photoId uint, newPhoto *models.Photo) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	err = DB.Where("id = ? and account_id = ?", photoId, accountId).First(photo).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	if newPhoto.Description != "" {
		photo.Description = newPhoto.Description
	}
	if newPhoto.Name != "" {
		photo.Name = newPhoto.Name
	}
	if newPhoto.Path != "" {
		photo.Path = newPhoto.Path
	}
	if newPhoto.GalleryId != 0 {
		photo.GalleryId = newPhoto.GalleryId
	}
	if newPhoto.Size != 0 {
		photo.Size = newPhoto.Size
	}

	err = DB.Save(photo).Error

	reactions, err := getReactionsByPhoto(*photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	photo.Reactions = reactions

	return
}

func DeletePhoto(accountId uint, photoId uint) (err error) {
	photo := &models.Photo{}
	err = DB.Where("id = ? and account_id = ?", photoId, accountId).First(photo).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	err = DB.Delete(photo).Error

	return
}

func getReactionsByPhoto(photo models.Photo) (reactions []models.Reaction, err error) {
	reactions = []models.Reaction{}
	err = DB.Model(&photo).Related(&reactions).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
