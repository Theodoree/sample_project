package app

import (
	"github.com/gin-gonic/gin"
	"github.com/Theodoree/sample_project/sample/mvc/pkg/errs"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		`msg`:  errs.GetMsg(errCode),
		"data": data,
	})

	return
}
