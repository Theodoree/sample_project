package app

import (
	"github.com/gin-gonic/gin"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		`msg`:errors.GetMsg(httpCode),
		"data": data,
	})

	return
}