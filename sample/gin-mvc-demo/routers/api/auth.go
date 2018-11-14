package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := errors.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password) //检测账户是否合法
		if isExist {
			token, err := util.GenerateToken(username, password) //返回签名后的Token
			if err != nil {
				code = errors.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token //设置token
				code = errors.SUCCESS
			}

		} else {
			code = errors.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errors.GetMsg(code),
		"data": data,
	})
}
