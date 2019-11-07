package main

import (
	"github.com/zhangtao25/mangostreetSerGin/routers"
	"github.com/zhangtao25/mangostreetSerGin/models"
)


func init() {
	models.Setup()
}


func main() {
	r := routers.InitRouter()

	// 默认启动的是 8080端口，也可以自己定义启动端口
	r.Run(":3000")
}