package routes

import (
	"errors"
	"fmt"
	"gallery/models"
	"gallery/services"

	"github.com/gin-gonic/gin"
)

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordObject struct {
	NewPassword string `json:"newPassword`
	OldPassword string `json:"oldPassword`
}

func Authentication(ctx *gin.Context) {
	cred := &Credential{}
	if err := ctx.BindJSON(cred); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(401, errors.New("Invalid email or Password"))
		return
	}

	token, err := services.Authenticate(cred.Email, cred.Password)
	if err != nil {
		ctx.AbortWithError(401, errors.New("Invalid email or password"))
		return
	}
	ctx.String(200, token)
}

func Registration(ctx *gin.Context) {
	cred := &Credential{}
	account := &models.Account{}
	if err := ctx.BindJSON(cred); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid email or Password"))
		return
	}

	account, err := services.Registration(cred.Email, cred.Password)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Invalid email or password"))
		return
	}
	ctx.JSON(200, account)
}

func GetAccount(ctx *gin.Context) {
	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	services.Logger.Infof("Get account information by id=[%d]", accountID)
	account, err := services.GetAccountByID(accountID.(uint))

	if err != nil {
		ctx.AbortWithError(404, errors.New("Account Not Found"))
		return
	}
	ctx.JSON(200, account)
}

func UpdateAccount(ctx *gin.Context) {
	account := &models.Account{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	if err := ctx.BindJSON(account); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid account JSON"))
		return
	}

	services.Logger.Infof("Update account information by id=[%d]", accountID)
	account, err := services.UpdateAccount(accountID.(uint), account)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Update failed"))
		return
	}

	ctx.JSON(200, account)
}

func ChangePassword(ctx *gin.Context) {
	changePasswordObject := &ChangePasswordObject{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	if err := ctx.BindJSON(changePasswordObject); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid account JSON"))
		return
	}

	services.Logger.Infof("Change account password by id=[%d]", accountID)
	account, err := services.ChangePassword(accountID.(uint), changePasswordObject.OldPassword, changePasswordObject.NewPassword)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Update failed"))
		return
	}

	ctx.JSON(200, account)
}

func SetAvatar(ctx *gin.Context) {
	account := &models.Account{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	fmt.Println(accountID)

	link := "uploads/" + fmt.Sprint(accountID.(uint)) + "/avatar/"
	path, _, err := saveUploadedFileToDirectory(ctx, link)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}

	services.Logger.Infof("Set account avatar by id=[%d]", accountID)
	account, err = services.SetAvatar(accountID.(uint), path)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Update failed"))
		return
	}

	ctx.JSON(200, account)
}

func DeleteAccount(ctx *gin.Context) {
	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	services.Logger.Infof("Delete account and related infomation by id=[%d]", accountID)
	err := services.DeleteAccount(accountID.(uint))

	if err != nil {
		ctx.AbortWithError(401, errors.New("Delete fail"))
		return
	}

	ctx.Status(200)
}
