package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Theodoree/sample_project/sample/mvc/router/auth"
	"github.com/Theodoree/sample_project/sample/mvc/router/User"
	"github.com/Theodoree/sample_project/sample/mvc/middleware/Jwt"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST(`/login`, auth.Login)
	r.POST(`/loginOut`, auth.LoginOut)
	r.POST(`/Create`, User.CreateNewUser)
	r.POST(`/Change`, User.ChangePassWord)

	v1 := r.Group(`/v1`, jwt.JWT())
	{
		v1.GET(`/HeeloWorld`, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				`Msg`: `Hello`,
			})
		})
	}

	return r
}
