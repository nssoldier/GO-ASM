package services

import (
	"errors"
	"fmt"
	"gallery/models"
)

func GetAccountByID(id uint) (account *models.Account, err error) {
	Logger.Debugf("Get account information by id=[%d]", id)
	account = &models.Account{}
	err = DB.First(account, id).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	photos, err := getPhotosByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}
	galleries, err := getGalleriesByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}
	reactions, err := getReactionsByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}

	account.Photos = photos
	account.Galleries = galleries
	account.Reactions = reactions

	return
}

func GetPublicAccountByID(id uint) (account *models.Account, err error) {
	Logger.Debugf("Get account information by id=[%d]", id)
	account = &models.Account{}
	err = DB.First(account, id).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	if account.Avatar == "" && account.Name == "" {
		err = errors.New("Can not find this public account")
		return
	}

	photos, err := getPhotosByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}
	galleries, err := getGalleriesByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}
	reactions, err := getReactionsByAccount(*account)
	if err != nil {
		fmt.Println(err)
		return
	}

	account.Photos = photos
	account.Galleries = galleries
	account.Reactions = reactions

	return
}

func GetAccountByEmail(email string) (account *models.Account, err error) {
	account = &models.Account{}
	err = DB.Where("email = ?", email).First(account).Error
	return
}

func UpdateAccount(id uint, newAccount *models.Account) (account *models.Account, err error) {
	account = &models.Account{}
	err = DB.Where("id = ?", id).First(account).Error

	if err != nil {
		return
	}

	if newAccount.Email != "" {
		account.Email = newAccount.Email
	}
	if newAccount.Password != "" {
		account.Password = newAccount.Password
	}
	if newAccount.Avatar != "" {
		account.Avatar = newAccount.Avatar
	}
	if newAccount.Name != "" {
		account.Name = newAccount.Name
	}

	err = DB.Save(account).Error

	return
}

func DeleteAccount(id uint) (err error) {
	Logger.Debugf("Delete account by id = [%d]", id)

	account := &models.Account{}
	err = DB.First(account, id).Error

	if err != nil {
		return
	}

	err = DB.Delete(account).Error

	return
}

func getGalleriesByAccount(account models.Account) (galleries []models.Gallery, err error) {
	galleries = []models.Gallery{}
	err = DB.Model(&account).Related(&galleries).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func getPhotosByAccount(account models.Account) (photos []models.Photo, err error) {
	photos = []models.Photo{}
	err = DB.Model(&account).Related(&photos).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func getReactionsByAccount(account models.Account) (reactions []models.Reaction, err error) {
	reactions = []models.Reaction{}
	err = DB.Model(&account).Related(&reactions).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
