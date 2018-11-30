package util

import (
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func Getpage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
