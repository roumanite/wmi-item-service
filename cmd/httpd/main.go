package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/item", func(c *gin.Context) {
		fmt.Println("Post item")
	})

	router.Run()
}
