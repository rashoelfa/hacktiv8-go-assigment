package routers

import (
	commentsController "mygram/app/controllers/comments"
	photosController "mygram/app/controllers/photos"
	socialMediasController "mygram/app/controllers/socialMedias"
	usersController "mygram/app/controllers/users"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", usersController.CreateUser)
		userRouter.POST("/login", usersController.UserLogin)
	}

	ProtectedUserRouter := r.Group("/users")
	{
		ProtectedUserRouter.Use(middlewares.Authentication())
		ProtectedUserRouter.PUT("/", usersController.UpdateUser)
		ProtectedUserRouter.DELETE("/", usersController.DeleteUser)
	}

	ProtectedPhotoRouter := r.Group("/photos")
	{
		ProtectedPhotoRouter.Use(middlewares.Authentication())
		ProtectedPhotoRouter.POST("/", photosController.CreatePhoto)
		ProtectedPhotoRouter.GET("/", photosController.GetAllPhotos)
		ProtectedPhotoRouter.PUT("/:id", photosController.UpdatePhoto)
		ProtectedPhotoRouter.DELETE("/:id", photosController.DeletePhoto)
	}

	ProtectedCommentsRouter := r.Group("/comments")
	{
		ProtectedCommentsRouter.Use(middlewares.Authentication())
		ProtectedCommentsRouter.POST("/", commentsController.CreateComments)
		ProtectedCommentsRouter.GET("/", commentsController.GetAllComments)
		ProtectedCommentsRouter.PUT("/:id", commentsController.UpdateComments)
		ProtectedCommentsRouter.DELETE("/:id", commentsController.DeleteComments)
	}

	ProtectedSocialMediasRouter := r.Group("/socialmedias")
	{
		ProtectedSocialMediasRouter.Use(middlewares.Authentication())
		ProtectedSocialMediasRouter.POST("/", socialMediasController.CreateSocialMedia)
		ProtectedSocialMediasRouter.GET("/", socialMediasController.GetAllSocialMedias)
		ProtectedSocialMediasRouter.PUT("/:id", socialMediasController.UpdateSocialMedia)
		ProtectedSocialMediasRouter.DELETE("/:id", socialMediasController.DeleteSocialMedia)
	}
}
