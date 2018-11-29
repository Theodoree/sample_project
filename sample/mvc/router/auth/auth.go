package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/gomysql"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/app"
	"github.com/Theodoree/sample_project/sample/mvc/models"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/errs"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/utils"
)

func Login(c *gin.Context) {
	Appc := app.Gin{c}
	Conn := gomysql.Db
	user := models.User{}
	if c.Bind(&user) != nil {
		Appc.Response(http.StatusOK, errs.INVALID_PARAMS, nil)
		return
	}
	Query := `select Salt,Pwd from user where Name= '%s'`
	rows, _, err := Conn.Query(Query, user.Name)
	if err != nil {
		Appc.Response(http.StatusOK, errs.ERROR_MYSQL_SELECT_FAIL, nil)
		return
	}
	Salt := rows[0].Str(0)
	Pwd := rows[0].Str(1)
	if utils.SaltToPassWord(user.Pwd, Salt) != Pwd {
		Appc.Response(http.StatusOK, errs.ERROR_PASSWORD_OR_USER, nil)
		return
	}
	token, err := utils.GenerateToken(user.Name, user.Pwd)
	if err != nil {
		Appc.Response(http.StatusOK, errs.ERROR_AUTH_TOKEN, nil)
		return
	}
	c.Set("token", token)
	data := map[string]interface{}{}
	data[`token`] = token
	Appc.Response(http.StatusOK, errs.SUCCESS, data)

}

func LoginOut(c *gin.Context) {
	c.Request.Header.Del(`Token`)
	c.JSON(http.StatusOK, gin.H{
		`ok`:  true,
		`msg`: "成功登出",
	})
	//这里应该有一个跳转页面 跳转到Index
}
