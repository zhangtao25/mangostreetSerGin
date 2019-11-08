package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangtao25/mangostreetSerGin/routers"
)


func main() {

	routers.C()





	r := gin.Default()

	r.GET("/user/:title", routers.GetNote)


	r.Run(":8080")
}