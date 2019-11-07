package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangtao25/mangostreetSerGin/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/add", v1.AddArticle)
	}

	return r
}
