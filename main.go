package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leon0399/cardyo-pdf/controllers"
	"rsc.io/quote"
)

func main() {
	router := gin.Default()

	router.GET("/quote/go", func(c *gin.Context) {
		c.String(http.StatusOK, quote.Go())
	})

	api := router.Group("/api/v1")

	{
		booklet := api.Group("/booklet")

		booklet.GET("/a5", controllers.GenerateA5Booklet)
	}

	router.Run()
}
