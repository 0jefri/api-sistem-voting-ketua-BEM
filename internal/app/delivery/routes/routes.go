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
	userController := controller.NewUserController(repoManager.UserService(), repoManager.AuthService())

	v1 := router.Group("/api/v1")
	{
		voting := v1.Group("/voting")
		{
			auth := voting.Group("/auth")
			{
				auth.POST("/register", userController.Registration)
				auth.POST("/login", userController.Login)
			}
		}
	}

	return router.Run()

}
