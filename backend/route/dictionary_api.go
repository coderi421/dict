package route

import (
	"dict/config"
	"dict/controller"
	"dict/middleware"
	"dict/workflow"

	"github.com/gin-gonic/gin"
)

func RouteDictionary(route *gin.Engine, userService workflow.UserService, dicService workflow.DictionaryService) {
	authService := config.NewServiceAuth()
	dictionaryController := controller.NewDictionaryController(dicService)
	dictionaryMiddleware := middleware.AuthMiddlewareUser(authService, userService)

	//root := route.Group("/")
	//root.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	api := route.Group("/api/v1/")
	api.GET("dictionary/search", dictionaryController.SearchDictionary)
	api.POST("dictionary/create", dictionaryController.CreateDictionary)
	api.POST("dictionary/update", dictionaryMiddleware, dictionaryController.UpdateDictionary)
	api.POST("dictionary/delete", dictionaryMiddleware, dictionaryController.DeleteDictionary)
}
