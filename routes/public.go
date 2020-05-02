package routes

import (
	"errors"
	"fmt"
	"gallery/models"
	"gallery/services"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Widths = [...]int{1920, 1600, 1280, 1024, 800, 256}

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

func GetPublicGalleryByName(ctx *gin.Context) {
	gallery := &models.Gallery{}

	galleryName := ctx.Param("name")

	gallery, err := services.GetPublicGalleryByName(galleryName)

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

	name, _ := ctx.GetQuery("name")

	if name == "" {
		galleries := &[]models.Gallery{}
		galleries, err := services.GetPublicGalleries()
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(400, errors.New("Get all galleries failed"))
			return
		}

		ctx.JSON(200, galleries)
	} else {
		gallery := &models.Gallery{}

		gallery, err := services.GetPublicGalleryByName(name)

		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(400, errors.New("Get gallery failed"))
			return
		}
		ctx.JSON(200, gallery)
	}

}

func DownloadPhotoByWidth(ctx *gin.Context) {
	photo := &models.Photo{}

	photoId := ctx.Param("id")
	id, err := strconv.ParseUint(photoId, 10, 64)

	photoWidth := ctx.Param("width")
	width, err := strconv.ParseInt(photoWidth, 10, 64)

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

	targetPath := ""
	fileName := ""

	for _, w := range Widths {
		fmt.Println(w)
		fmt.Println(int(width))
		if w == int(width) {
			if w == 1920 {
				targetPath = photo.W1920Path
			}
			if w == 1600 {
				targetPath = photo.W1600Path
			}
			if w == 1280 {
				targetPath = photo.W1280Path
			}
			if w == 1024 {
				targetPath = photo.W1024Path
			}
			if w == 800 {
				targetPath = photo.W800Path
			}
			if w == 256 {
				targetPath = photo.W256Path
			}
			break
		}
	}

	if targetPath == "" {
		ctx.AbortWithError(400, errors.New("can not find this width"))
		return
	}

	fileInfo, err := os.Stat(targetPath)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get photo failed"))
		return
	}

	fileName = fileInfo.Name()
	ctx.FileAttachment(targetPath, fileName)

}
