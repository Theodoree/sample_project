package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
	"github.com/Unknwon/com"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"strconv"
	"github.com/astaxie/beego/validation"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := errors.SUCCESS

	data["lists"] = models.GetTags(util.Getpage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errors.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query(`name`)
	state, _ := strconv.Atoi(c.DefaultQuery(`state`, `0`))
	createdBy := c.Query(`created_by`)

	valid := validation.Validation{}
	valid.Required(name, `name`).Message(`名称不能为空`)
	valid.MaxSize(name, 80, `name`).Message(`名称最长为二十个函数或者八十个字符`)
	valid.Required(createdBy, `created_by`).Message(`创建人不能为空`)
	valid.MaxSize(createdBy, 80, `created_by`).Message(`创建人不能为空`)
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errors.INVALID_PARAMS
	if ! valid.HasErrors() { //验证是否有错误
		if !models.ExistTagByName(name) { //检测是否存在
			code = errors.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = errors.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)

		}
	}
	c.JSON(http.StatusOK, gin.H{
	`code`: code,
	`msg`:  errors.GetMsg(code),
	`data`: make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name :=c.Query(`name`)
	modifiedBy := c.Query(`modified_by`)

	valid := validation.Validation{}

	var state = -1
	if arg :=c.Query(`state`);arg != ""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,`state`).Message(`状态只允许0或1`)
	}

	valid.Required(id,`id`).Message(`ID不能为空`)
	valid.Required(modifiedBy,`modified_by`).Message(`状态不能为空`)
	valid.MaxSize(modifiedBy,80,`modified_by`).Message(`最长为二十个汉字或八十个字符`)
	valid.MaxSize(name,80,`name`).Message(`最长为二十个汉字或八十个字符`)

	code:= errors.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errors.SUCCESS
		if models.ExistTagByID(id){
			data :=make(map[string]interface{})
			data[`modified_by`] = modifiedBy
			if name !=""{
				data[`name`] = name
			}
			if state != -1 {
				data[`state`] = state
			}
			models.EditTag(id,data)
		}else{
			code = errors.ERROR_NOT_EXIST_TAG //没有找到
		}
	}else{
		for _,err:=range valid.Errors{
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		`code`: code,
		`msg`:  errors.GetMsg(code),
		`data`: make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id :=com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errors.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = errors.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = errors.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" :  errors.GetMsg(code),
		"data" : make(map[string]string),
	})
}