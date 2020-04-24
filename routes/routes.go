package routes

import (
	"gallery/middlewares"

	"github.com/gin-gonic/gin"
)

func Create() (g *gin.Engine) {

	g = gin.Default()
	// g.Use(static.Serve("/image", static.LocalFile("/image", true)))
	v1 := g.Group("/v1")
	{
		v1.POST("/registration", Registration)
		v1.POST("/authentication", Authentication)

		account := v1.Group("/account")
		{
			// Use authentication middlewares
			account.Use(middlewares.RequireAuthentication())
			// add account handlers
			account.GET("", GetAccount)
			account.PUT("", UpdateAccount)
			account.DELETE("", DeleteAccount)
		}

		gallery := v1.Group("/gallery")
		{
			// Use authentication middlewares
			gallery.Use(middlewares.RequireAuthentication())
			// add gallery handlers
			gallery.POST("", CreateGallery)
			gallery.GET("", GetAllGalleries)
			gallery.GET("/:id", GetGallery)
			gallery.PUT("/:id", UpdateGallery)
			gallery.DELETE("/:id", DeleteGallery)
		}
		photo := v1.Group("/photo")
		{
			// Use authentication middlewares
			photo.Use(middlewares.RequireAuthentication())
			// add photo handlers
			photo.POST("", CreatePhoto)
			photo.GET("/:id", GetPhoto)
			photo.PUT("/:id", UpdatePhoto)
			photo.DELETE("/:id", DeletePhoto)
			photo.POST("/:id/reaction", CreateReaction)
			photo.DELETE("/:id/reaction", DeleteReaction)
		}
		public := v1.Group("/public")
		{
			// add public handlers

			public.GET("/gallery", GetPublicGalleries)
			public.GET("/gallery/:id", GetPublicGallery)
			public.GET("/photo/:id", GetPublicPhoto)
			public.GET("/account/:id", GetPublicAccount)

		}
	}

	return
}
