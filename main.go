package main

import (
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leon0399/cardyo-pdf/controllers"
	"rsc.io/quote"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	router := gin.Default()

	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	router.GET("/quote/go", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Go())
	})

	api := router.Group("/api/v1")

	{
		booklet := api.Group("/booklet")

		booklet.GET("/a5", controllers.GenerateBookletA5Api)
	}

	router.Run()
}
