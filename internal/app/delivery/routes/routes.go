package routes

import (
	"github.com/api-voting/config"
	"github.com/api-voting/internal/app/delivery/controller"
	"github.com/api-voting/internal/app/delivery/manager"
	"github.com/api-voting/internal/app/delivery/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// User Controller
	userController := controller.NewUserController(repoManager.UserService())

	v1 := router.Group("/api/v1")
	{
		voting := v1.Group("/voting")
		{
			auth := voting.Group("/auth")
			{
				auth.POST("/register", userController.Registration)
				// auth.POST("/login", userController.Login)
			}

			// users := sakupay.Group("/users", middleware.AuthMiddleware())
			// {
			// 	users.GET("/", userController.FindAllUsers)
			// 	users.GET("/:id", userController.FindUser)
			// 	users.PUT("/:id", userController.UpdateUser)
			// 	users.POST("/:id/upload", userController.UploadPicture)
			// 	users.GET("/:id/download", userController.DownloadPicture)
			// 	users.DELETE("/:id", userController.DeleteUser)
			// }
		}
	}

	return router.Run()

}
