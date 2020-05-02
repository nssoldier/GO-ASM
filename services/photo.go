package services

import (
	"fmt"
	"gallery/models"
	"os"
)

func CreatePhoto(accountId uint, newPhoto *models.Photo) (photo *models.Photo, err error) {
	photo = newPhoto
	photo.AccountId = accountId

	resizePaths, err := ResizeAll(photo.Path)
	if resizePaths[1920] != "" {
		photo.W1920Path = resizePaths[1920]
	}
	if resizePaths[1600] != "" {
		photo.W1600Path = resizePaths[1600]
	}
	if resizePaths[1280] != "" {
		photo.W1280Path = resizePaths[1280]
	}
	if resizePaths[1024] != "" {
		photo.W1024Path = resizePaths[1024]
	}
	if resizePaths[800] != "" {
		photo.W800Path = resizePaths[800]
	}
	if resizePaths[256] != "" {
		photo.W256Path = resizePaths[256]
	}
	err = DB.Create(&photo).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetPhotoById(photoId uint, accountId uint) (photo *models.Photo, err error) {

	photo = &models.Photo{}
	err = DB.Where("id = ? and account_id = ?", photoId, accountId).First(photo, photoId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	reactions, err := GetReactionsByPhoto(*photo)
	if err != nil {
		fmt.Println(err)
		return
	}
	photo.Reactions = reactions

	return
}

func GetPublicPhotoById(photoId uint) (photo *models.Photo, err error) {

	photo = &models.Photo{}
	err = DB.Where("id = ?", photoId).First(photo, photoId).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = GetPublicGalleryById(photo.GalleryId)

	if err != nil {
		fmt.Println(err)
		return
	}

	reactions, err := GetReactionsByPhoto(*photo)
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

		fmt.Println(newPhoto.Path)
		resizePaths, err1 := ResizeAll(newPhoto.Path)
		fmt.Println(resizePaths)
		if resizePaths[1920] != "" {
			photo.W1920Path = resizePaths[1920]
		} else {
			DeleteFile(photo.W1920Path)
			photo.W1920Path = ""
		}
		if resizePaths[1600] != "" {
			photo.W1600Path = resizePaths[1600]
		} else {
			DeleteFile(photo.W1600Path)
			photo.W1600Path = ""
		}
		if resizePaths[1280] != "" {
			photo.W1280Path = resizePaths[1280]
		} else {
			DeleteFile(photo.W1280Path)
			photo.W1280Path = ""
		}
		if resizePaths[1024] != "" {
			photo.W1024Path = resizePaths[1024]
		} else {
			DeleteFile(photo.W1024Path)
			photo.W1024Path = ""
		}
		if resizePaths[800] != "" {
			photo.W800Path = resizePaths[800]
		} else {
			DeleteFile(photo.W800Path)
			photo.W800Path = ""
		}
		if resizePaths[256] != "" {
			photo.W256Path = resizePaths[256]
		} else {
			DeleteFile(photo.W256Path)
			photo.W256Path = ""
		}

		if err1 != nil {
			fmt.Println(err)
			return
		}
		photo.Path = newPhoto.Path
	}
	if newPhoto.GalleryId != 0 {
		photo.GalleryId = newPhoto.GalleryId
	}
	if newPhoto.Size != 0 {
		photo.Size = newPhoto.Size
	}

	err = DB.Save(photo).Error

	reactions, err := GetReactionsByPhoto(*photo)
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

	if photo.W1920Path != "" {
		DeleteFile(photo.W1920Path)
	}
	if photo.W1600Path != "" {
		DeleteFile(photo.W1600Path)
	}
	if photo.W1280Path != "" {
		DeleteFile(photo.W1280Path)
	}
	if photo.W1024Path != "" {
		DeleteFile(photo.W1024Path)
	}
	if photo.W800Path != "" {
		DeleteFile(photo.W800Path)
	}
	if photo.W256Path != "" {
		DeleteFile(photo.W256Path)
	}

	err = DB.Delete(photo).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetReactionsByPhoto(photo models.Photo) (reactions []models.Reaction, err error) {
	reactions = []models.Reaction{}
	err = DB.Model(&photo).Related(&reactions).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}
