package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Theodoree/sample_project/sample/mvc/pkg/errs"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = errs.SUCCESS
		token := c.GetHeade

		r(`token`)
		if token == "" {
			code = errs.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = errs.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errs.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errs.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errs.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
