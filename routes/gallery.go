package routes

import (
	"errors"
	"fmt"
	"gallery/models"
	"gallery/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGallery(ctx *gin.Context) {
	gallery := &models.Gallery{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	if err := ctx.BindJSON(gallery); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid gallery JSON"))
		return
	}

	gallery, err := services.CreateGallery(accountID.(uint), gallery)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Create gallery failed"))
		return
	}

	ctx.JSON(200, gallery)
}

func GetAllGalleries(ctx *gin.Context) {
	galleries := &[]models.Gallery{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}
	galleries, err := services.GetAllGalleries(accountID.(uint))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get all galleries failed"))
		return
	}

	ctx.JSON(200, galleries)
}

func GetGallery(ctx *gin.Context) {
	gallery := &models.Gallery{}

	galleryId := ctx.Param("id")

	id, err := strconv.ParseUint(galleryId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	gallery, err = services.GetGalleryById(uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get gallery failed"))
		return
	}
	ctx.JSON(200, gallery)
}

func UpdateGallery(ctx *gin.Context) {
	gallery := &models.Gallery{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	galleryId := ctx.Param("id")
	id, err := strconv.ParseUint(galleryId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	if err := ctx.BindJSON(gallery); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid gallery JSON"))
		return
	}

	gallery, err = services.UpdateGallery(accountID.(uint), uint(id), gallery)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get gallery failed"))
		return
	}
	ctx.JSON(200, gallery)
}

func DeleteGallery(ctx *gin.Context) {

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	galleryId := ctx.Param("id")
	id, err := strconv.ParseUint(galleryId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	err = services.DeleteGallery(accountID.(uint), uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get gallery failed"))
		return
	}

	ctx.Status(200)
}
