package routes

import (
	"errors"
	"fmt"
	"gallery/models"
	"gallery/services"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	photo := &models.Photo{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	galleryId := ctx.PostForm("galleryId")
	id, err := strconv.ParseUint(galleryId, 10, 64)

	if err != nil {
		ctx.AbortWithError(400, errors.New("Can not convert gallery id"))
		return
	}

	photo.GalleryId = uint(id)
	photo.Name = ctx.PostForm("name")
	photo.Description = ctx.PostForm("description")

	path, size, err := saveUploadedFileToDirectory(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}

	photo.Path = path
	photo.Size = size

	photo, err = services.CreatePhoto(accountID.(uint), photo)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Create photo failed"))
		return
	}

	ctx.JSON(200, photo)
}

func GetPhoto(ctx *gin.Context) {
	photo := &models.Photo{}

	photoId := ctx.Param("id")

	id, err := strconv.ParseUint(photoId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	photo, err = services.GetPhotoById(uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get photo failed"))
		return
	}
	ctx.JSON(200, photo)
}

func UpdatePhoto(ctx *gin.Context) {
	photo := &models.Photo{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	photoId := ctx.Param("id")
	id, err := strconv.ParseUint(photoId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}
	galleryId := ctx.PostForm("galleryId")
	id, err = strconv.ParseUint(galleryId, 10, 64)

	if err != nil {
		ctx.AbortWithError(400, errors.New("Can not convert gallery id"))
		return
	}

	photo.GalleryId = uint(id)
	photo.Name = ctx.PostForm("name")
	photo.Description = ctx.PostForm("description")

	path, size, err := saveUploadedFileToDirectory(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Can not save uploaded file"))
		return
	}

	photo.Path = path
	photo.Size = size

	photo, err = services.UpdatePhoto(accountID.(uint), uint(id), photo)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get photo failed"))
		return
	}
	ctx.JSON(200, photo)
}

func DeletePhoto(ctx *gin.Context) {

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	photoId := ctx.Param("id")
	id, err := strconv.ParseUint(photoId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	err = services.DeleteGallery(accountID.(uint), uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get photo failed"))
		return
	}

	ctx.Status(200)
}

func CreateReaction(ctx *gin.Context) {
	reaction := &models.Reaction{}

	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	photoId := ctx.Param("id")

	id, err := strconv.ParseUint(photoId, 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	reaction.PhotoId = uint(id)

	if err := ctx.BindJSON(reaction); err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Invalid reaction JSON"))
		return
	}

	reaction, err = services.CreateReaction(accountID.(uint), reaction)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Create reaction failed"))
		return
	}

	ctx.JSON(200, reaction)
}

func DeleteReaction(ctx *gin.Context) {
	accountID, exists := ctx.Get("account_id")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	reactionId := ctx.Param("id")
	id, err := strconv.ParseUint(reactionId, 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Convert param failed"))
		return
	}

	err = services.DeleteReaction(accountID.(uint), uint(id))

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithError(400, errors.New("Get reaction failed"))
		return
	}

	ctx.Status(200)
}

func saveUploadedFileToDirectory(c *gin.Context) (path string, size int64, err error) {

	file, err := c.FormFile("image")
	filename := filepath.Base(file.Filename)
	path = "image/" + filename
	size = file.Size
	if err = c.SaveUploadedFile(file, path); err != nil {
		fmt.Println(err)
		return
	}
	return
}
