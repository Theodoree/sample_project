package User

import (
	"github.com/gin-gonic/gin"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/gomysql"
	"github.com/Theodoree/sample_project/sample/mvc/models"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/app"
	"net/http"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/errs"

	"github.com/Theodoree/sample_project/sample/mvc/pkg/utils"
)

func CreateNewUser(c *gin.Context) {
	Appc := app.Gin{c}
	Conn := gomysql.Db
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		Appc.Response(http.StatusOK, errs.ERROR_PASSWORD_OR_USER, nil)
		return
	}
	Query := `INSERT INTO user(Name,Email,Phone,Pwd,Salt,InsertTime) VALUES ('%s','%s','%s','%s','%s','%s')`
	user.Insert = utils.Get_Now()
	user.Salt = utils.CreateSalt()
	user.Pwd = utils.SaltToPassWord(user.Pwd, user.Pwd)
	_, _, err = Conn.Query(Query, Conn.Escape(user.Name), Conn.Escape(user.Email), Conn.Escape(user.Phone),
		Conn.Escape(user.Pwd), Conn.Escape(user.Salt), Conn.Escape(user.Insert), )
	if err != nil {
		Appc.Response(http.StatusOK, errs.ERROR_MYSQL_INSERT_FAIL, nil)
		return
	}

	Appc.Response(http.StatusOK, errs.SUCCESS, nil)
}

func ChangePassWord(c *gin.Context) {
	Appc := app.Gin{c}
	Conn := gomysql.Db
	user := models.User{}
	if c.Bind(&user) != nil {
		Appc.Response(http.StatusOK, errs.ERROR_PASSWORD_OR_USER, nil)
		return
	}

	NewPassWord := c.PostForm(`NewPwd`)
	Query := `select Salt,Pwd from user where Name= '%s'`
	UpQuery := `update user set Pwd='%s' where Name= '%s'`
	rows, _, err := Conn.Query(Query, user.Name)
	if err != nil {
		Appc.Response(http.StatusOK, errs.ERROR_MYSQL_SELECT_FAIL, nil)
		return
	}
	Salt := rows[0].Str(0)
	Pwd := rows[0].Str(1)
	if utils.SaltToPassWord(user.Pwd, Salt) == Pwd {
		NewPassWord = utils.SaltToPassWord(NewPassWord, Salt)
		_, _, err := Conn.Query(UpQuery, NewPassWord, user.Name)
		if err != nil {
			Appc.Response(http.StatusOK, errs.ERROR_MYSQL_UPDATE_FAIL, nil)
			return
		}
	} else {
		Appc.Response(http.StatusOK, errs.ERROR_PASSWORD, nil)
		return
	}
	Appc.Response(http.StatusOK, errs.SUCCESS, nil)
}
