package routes

import (
	"errors"
	"fmt"
	"gallery/models"
	"gallery/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPublicGallery(ctx *gin.Context) {
	gallery := &models.Gallery{}

	galleryId := ctx.Param("id")

	id, err := strconv.ParseUint(galleryId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	gallery, err = services.GetPublicGalleryById(uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get gallery failed"))
		return
	}
	ctx.JSON(200, gallery)
}
func GetPublicPhoto(ctx *gin.Context) {
	photo := &models.Photo{}

	photoId := ctx.Param("id")

	id, err := strconv.ParseUint(photoId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	photo, err = services.GetPublicPhotoById(uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get photo failed"))
		return
	}
	ctx.JSON(200, photo)
}
func GetPublicAccount(ctx *gin.Context) {
	account := &models.Account{}

	accountId := ctx.Param("id")

	id, err := strconv.ParseUint(accountId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	account, err = services.GetPublicAccountByID(uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get account failed"))
		return
	}
	ctx.JSON(200, account)
}
func GetPublicGalleries(ctx *gin.Context) {
	galleries := &[]models.Gallery{}

	galleries, err := services.GetPublicGalleries()

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get all galleries failed"))
		return
	}

	ctx.JSON(200, galleries)
}
