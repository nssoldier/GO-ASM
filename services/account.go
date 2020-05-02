package services

import (
	"errors"
	"fmt"
	"gallery/models"

	"golang.org/x/crypto/bcrypt"
)

func GetAccountByID(id uint) (account *models.Account, err error) {
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
	account = &models.Account{}
	err = DB.Where("id = ?", id).First(account).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account.Avatar)
	fmt.Println(account.Name)
	fmt.Println(account.Address)
	fmt.Println(account.Phone)
	if account.Avatar == "" || account.Name == "" || account.Address == "" || account.Phone == "" {
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
	if newAccount.Address != "" {
		account.Address = newAccount.Address
	}
	if newAccount.Phone != "" {
		account.Phone = newAccount.Phone
	}
	if newAccount.Name != "" {
		account.Name = newAccount.Name
	}

	err = DB.Save(account).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func ChangePassword(id uint, oldPasword string, newPassword string) (account *models.Account, err error) {
	account = &models.Account{}
	err = DB.Where("id = ?", id).First(account).Error

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(oldPasword))

	if err != nil {
		fmt.Println(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	account.Password = string(hashedPassword)
	err = DB.Save(account).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func SetAvatar(id uint, path string) (account *models.Account, err error) {
	account = &models.Account{}
	err = DB.Where("id = ?", id).First(account).Error

	if err != nil {
		return
	}

	if path != "" {
		account.Avatar = path
	}

	err = DB.Save(account).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DeleteAccount(id uint) (err error) {
	Logger.Debugf("Delete account by id = [%d]", id)

	account := &models.Account{}
	err = DB.Where("id = ?", id).First(account, id).Error

	if err != nil {
		return
	}

	if account.Avatar != "" {
		DeleteFile(account.Avatar)
	}

	galleries := &[]models.Gallery{}
	err = DB.Where("account_id = ?", account.Id).Find(galleries).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.Delete(galleries)

	photos := &[]models.Photo{}
	err = DB.Where("account_id = ?", account.Id).Find(photos).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	for index := 0; index < len(*photos); index++ {
		DeletePhoto(account.Id, (*photos)[index].Id)
	}

	reactions := &[]models.Reaction{}
	err = DB.Where("account_id = ?", account.Id).Find(reactions).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	DB.Delete(reactions)

	err = DB.Delete(account).Error

	if err != nil {
		fmt.Println(err)
		return
	}

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
