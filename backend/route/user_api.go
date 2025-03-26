package route

import (
	"dict/config"
	"dict/controller"
	"dict/middleware"
	"dict/workflow"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, service workflow.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)
	userMiddleware := middleware.AuthMiddlewareUser(authService, service) // middl.AuthMiddlewareManager(authService, workflow)
	root := route.Group("/")
	root.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := root.Group("/api/v1/")
	api.POST("user/login", userController.Login)
	api.POST("update-account", userMiddleware, userController.UpdateProfile)
}
