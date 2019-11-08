package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangtao25/mangostreetSerGin/models"
	"github.com/zhangtao25/mangostreetSerGin/routers"
)


func main() {
	models.Setup()

	r := gin.Default()
	r.GET("/user/:title", routers.GetNote)

	r.Run(":8080")
}