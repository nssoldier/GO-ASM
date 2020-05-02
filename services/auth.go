package services

import (
	"errors"
	"fmt"
	"gallery/models"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(email string, password string) (tokenStr string, err error) {
	account, err := GetAccountByEmail(email)
	if err != nil {
		return

	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	tokenStr, err = CreateToken(account.Id)
	return
}

func Registration(email string, password string) (account *models.Account, err error) {
	account = &models.Account{}
	account.Email = email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	account.Password = string(hashedPassword)

	err = DB.Create(&account).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
