package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangtao25/mangostreetSerGin/pkg/app"
	"github.com/zhangtao25/mangostreetSerGin/pkg/e"
	"github.com/zhangtao25/mangostreetSerGin/service/user_service"
	"net/http"
)


type AuthUsersByVerificationCodeForm struct {
	Username       string `form:"username" valid:"Required;MaxSize(100)"`
	Vcode          string `form:"vcode"    valid:"Required;MaxSize(255)"`
}


func AuthUsersByVerificationCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = AuthUsersByVerificationCodeForm{}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	fmt.Print(form.Username,654)

	userService := user_service.User{
		Username:form.Username,
		Vcode:form.Vcode,
	}

	token, err := userService.Get()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}