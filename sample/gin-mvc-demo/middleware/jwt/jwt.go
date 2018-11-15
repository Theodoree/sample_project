package jwt


import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = errors.SUCCESS
		token := c.GetHeader(`token`)
		if token == "" {
			code = errors.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = errors.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errors.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errors.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : errors.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}