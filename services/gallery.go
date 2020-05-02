package services

import (
	"errors"
	"fmt"
	"gallery/models"
)

func CreateGallery(accountId uint, newGallery *models.Gallery) (gallery *models.Gallery, err error) {
	gallery = newGallery
	gallery.AccountId = accountId

	err = DB.Create(&gallery).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetAllGalleries(accountId uint) (galleries *[]models.Gallery, err error) {
	account := &models.Account{}
	galleries = &[]models.Gallery{}
	err = DB.First(account, accountId).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	err = DB.Model(&account).Related(galleries).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, gallery := range *galleries {
		photos := []models.Photo{}
		photos, err = getPhotosByGallery(gallery)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(photos) > 3 {
			(*galleries)[index].Photos = photos[0:2]
		} else {
			(*galleries)[index].Photos = photos
		}
	}

	return
}

func GetPublicGalleries() (galleries *[]models.Gallery, err error) {
	galleries = &[]models.Gallery{}
	err = DB.Where("visibility = ?", 1).Find(galleries).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	for index, gallery := range *galleries {
		photos := []models.Photo{}
		photos, err = getPhotosByGallery(gallery)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(photos) > 3 {
			(*galleries)[index].Photos = photos[0:2]
		} else {
			(*galleries)[index].Photos = photos
		}
	}

	return
}

func GetGalleryById(accountId uint, galleryId uint) (gallery *models.Gallery, err error) {

	gallery = &models.Gallery{}
	err = DB.Where("id = ? and account_id = ?", galleryId, accountId).First(gallery, galleryId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	photos, err := getPhotosByGallery(*gallery)
	if err != nil {
		fmt.Println(err)
		return
	}
	gallery.Photos = photos

	return
}

func GetPublicGalleryById(galleryId uint) (gallery *models.Gallery, err error) {

	gallery = &models.Gallery{}
	err = DB.Where("id = ?", galleryId).First(gallery, galleryId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	if !gallery.Visibility {
		err = errors.New("Can not find this public gallery")
		return
	}

	photos, err := getPhotosByGallery(*gallery)
	if err != nil {
		fmt.Println(err)
		return
	}
	gallery.Photos = photos

	return
}

func GetPublicGalleryByName(galleryName string) (gallery *models.Gallery, err error) {

	gallery = &models.Gallery{}
	err = DB.Where("name = ?", galleryName).First(gallery).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	if !gallery.Visibility {
		err = errors.New("Can not find this public gallery")
		return
	}

	photos, err := getPhotosByGallery(*gallery)
	if err != nil {
		fmt.Println(err)
		return
	}
	gallery.Photos = photos

	return
}

func UpdateGallery(accountId uint, galleryId uint, newGallery *models.Gallery) (gallery *models.Gallery, err error) {
	gallery = &models.Gallery{}
	err = DB.Where("id = ? and account_id = ?", galleryId, accountId).First(gallery).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	if newGallery.Name != "" {
		gallery.Name = newGallery.Name
	}
	if newGallery.Brief != "" {
		gallery.Brief = newGallery.Brief
	}

	err = DB.Save(gallery).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	photos, err := getPhotosByGallery(*gallery)
	if err != nil {
		fmt.Println(err)
		return
	}
	gallery.Photos = photos
	return
}

func PublicGallery(accountId uint, galleryId uint) (err error) {
	gallery := &models.Gallery{}
	err = DB.Where("id = ? and account_id = ?", galleryId, accountId).First(gallery).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	if !gallery.Visibility {
		gallery.Visibility = true
	}

	err = DB.Save(gallery).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DeleteGallery(accountId uint, galleryId uint) (err error) {
	gallery := &models.Gallery{}
	err = DB.Where("id = ? and account_id = ?", galleryId, accountId).First(gallery).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	err = DB.Delete(gallery).Error

	return
}

func getPhotosByGallery(gallery models.Gallery) (photos []models.Photo, err error) {
	photos = []models.Photo{}
	err = DB.Where("gallery_id = ?", gallery.Id).Find(&photos).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	for index, photo := range photos {
		reactions := []models.Reaction{}
		reactions, err = GetReactionsByPhoto(photo)
		if err != nil {
			fmt.Println(err)
			return
		}

		photos[index].ReactionCount = len(reactions)
	}

	return
}
