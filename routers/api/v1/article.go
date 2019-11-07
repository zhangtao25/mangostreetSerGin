package v1

import (
	"github.com/zhangtao25/mangostreetSerGin/pkg/app"
	"github.com/zhangtao25/mangostreetSerGin/pkg/e"
	"github.com/zhangtao25/mangostreetSerGin/service/article_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	articleService := article_service.Article{
		TagID:         111,
		Title:         "grdgdfg",
		Desc:          "sdgdfg",
		Content:       "rgdrgdfgdf",
		CoverImageUrl: "ryherhd",
		State:         0,
	}

	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}