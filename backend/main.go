package main

import (
	"context"
	"dict/helper"
	"dict/middleware"
	"dict/repository"
	"dict/route"
	"dict/workflow"
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	helper.InitLogger()
	err := godotenv.Load()
	if err != nil {
		helper.Log.Error(context.Background(), "Error getting env, not comming through %v", err)
	} else {
		helper.Log.Info(context.Background(), "We are getting the env values")
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_API"),
	})
	if err != nil {
		helper.Log.Error(context.Background(), "sentry.Init: %s", err)
	}

	gin.SetMode(gin.DebugMode)

	userRepo := repository.NewUserRepository()
	userService := workflow.NewUserService(userRepo)

	hotKeyRepo := repository.NewSearchHotKeywordRepository(100)
	hotKeyRepo.KeywordCache.StartSync(5 * 60 * time.Second)
	//hotKeyRepo.KeywordCache.StartSync(5 * time.Second)
	hotKeyService := workflow.NewSearchHotKeywordService(hotKeyRepo)

	dictRepo := repository.NewDictionaryRepository()
	dictService := workflow.NewDictionaryService(dictRepo)

	categoryRepo := repository.NewCategoryRepository()
	categoryService := workflow.NewCategoryService(categoryRepo)

	secureMiddleware := middleware.SecureMiddleware()

	router := gin.Default()
	//root := router.Group("/")
	//root.GET("ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//api := root.Group("/api/v1/")

	router.Use(secureMiddleware)
	router.Use(sentrygin.New(sentrygin.Options{}))

	route.RouteUser(router, userService)
	route.RouteSearchHotKey(router, hotKeyService)
	route.RouteDictionary(router, userService, dictService, hotKeyService)
	route.RouteCategory(router, userService, categoryService)

	//router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.Log.Error(context.Background(), fmt.Sprintf("listen error: %v", err))
			// project start failed
			panic(fmt.Sprintf("listen error: %v", err))
		}
	}()

	helper.Log.Info(context.Background(), fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	helper.Log.Info(context.Background(), "Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		helper.Log.Error(ctx, "Server forced to shutdown: %v", err)
	}
	helper.Log.Info(ctx, "Server exiting")
}
