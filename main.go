package main

import (
	"model"

	"github.com/gin-gonic/gin"
	"github.com/wells7778/buDai/internal/controller"
)

func main() {
	router := gin.Default()
	model.Init()
	router.GET("/roles", controller.Index)
	router.GET("/user/:id", controller.Show)
	router.POST("/user", controller.Create)
	router.PUT("/user/:id", controller.Update)
	router.DELETE("/user/:id", controller.Destroy)
	router.Run(":8080")
}
