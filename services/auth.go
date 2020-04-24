package services

import (
	"errors"
	"fmt"
	"gallery/models"
)

func Authenticate(email string, password string) (tokenStr string, err error) {
	account, err := GetAccountByEmail(email)
	if err != nil {
		return

	}
	fmt.Println(account)
	if account.Password != password {
		return "", errors.New("Invalid email or password")
	}
	tokenStr, err = CreateToken(account.Id)
	return
}

func Registration(email string, password string) (account *models.Account, err error) {
	account = &models.Account{}
	account.Email = email
	account.Password = password

	err = DB.Create(&account).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
