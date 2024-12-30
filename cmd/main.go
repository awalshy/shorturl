package main

import (
	"github.com/awalshy/shorturl/pkg/shortener"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	shortener := shortener.NewShortener()
	shortener.RegisterRoutes(router)

	router.Run(":8080")
}