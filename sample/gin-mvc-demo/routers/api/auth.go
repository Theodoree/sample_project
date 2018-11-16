package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/app"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{c}

	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(errors.SUCCESS, errors.INVALID_PARAMS, data)
		return
	}

	isExist, err := models.CheckAuth(username, password) //检测账户是否合法
	if !isExist {
		appG.Response(errors.SUCCESS, errors.ERROR_AUTH, data)
		return
	}
	token, err := util.GenerateToken(username, password) //返回签名后的Token
	if err != nil {
		appG.Response(errors.SUCCESS, errors.ERROR_AUTH_TOKEN, data)
		return
	}
	data["token"] = token //设置token

	appG.Response(errors.SUCCESS, errors.SUCCESS, data)
}
