package v1

import (
	"net/http"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/app"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/export"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/service/tag_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	AppG := app.Gin{c}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.Getpage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	data, err := tagService.GetAll()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	AppG.Response(http.StatusOK, errors.SUCCESS, data)
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query(`name`)
	state, _ := strconv.Atoi(c.DefaultQuery(`state`, `0`))
	createdBy := c.Query(`created_by`)
	valid := validation.Validation{}
	AppG := app.Gin{c}
	valid.Required(name, `name`).Message(`名称不能为空`)
	valid.MaxSize(name, 80, `name`).Message(`名称最长为二十个函数或者八十个字符`)
	valid.Required(createdBy, `created_by`).Message(`创建人不能为空`)
	valid.MaxSize(createdBy, 80, `created_by`).Message(`创建人不能为空`)
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{Name: name, State: state, CreatedBy: createdBy}
	exists, err := tagService.ExistByName()
	if err != nil { //获取标签失败
		AppG.Response(http.StatusOK, errors.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if exists { //标签已经存在
		AppG.Response(http.StatusOK, errors.ERROR_EXIST_TAG, nil)
		return
	}
	err = tagService.Add()
	if err != nil { //添加失败
		AppG.Response(http.StatusOK, errors.ERROR_ADD_TAG_FAIL, nil)
		return
	}
	models.AddTag(name, state, createdBy)

	AppG.Response(http.StatusOK, errors.SUCCESS, nil)
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query(`name`)
	modifiedBy := c.Query(`modified_by`)
	AppG := app.Gin{c}
	valid := validation.Validation{}

	var state = -1
	if arg := c.Query(`state`); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, `state`).Message(`状态只允许0或1`)
	}
	valid.Required(id, `id`).Message(`ID不能为空`)
	valid.Required(modifiedBy, `modified_by`).Message(`状态不能为空`)
	valid.MaxSize(modifiedBy, 80, `modified_by`).Message(`最长为八十个字符`)
	valid.MaxSize(name, 80, `name`).Message(`最长为八十个字符`)
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}
	tagService := tag_service.Tag{ID: id, Name: name, State: state, ModifiedBy: modifiedBy}
	exists, err := tagService.ExistByID()
	if err != nil { //获取标签失败
		AppG.Response(http.StatusOK, errors.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists { //没有获取到Tag
		AppG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	err = tagService.Edit()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_EDIT_TAG_FAIL, nil)
		return
	}
	AppG.Response(http.StatusOK, errors.SUCCESS, nil)
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	appG := app.Gin{c}
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}
	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, errors.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusOK, errors.ERROR_DELETE_TAG_FAIL, nil)
	}

	appG.Response(http.StatusOK, errors.SUCCESS, nil)
}

// @Summary 导出文章标签
// @Produce  json
// @Param name post string false "Name"
// @Param state post int false "State"
// @Success 200 {string} json "{"code":200,"data":{"export_save_url":"export/abc.xlsx", "export_url": "http://..."},"msg":"ok"}"
// @Router /api/v1/tags/export [post]
func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:  name,
		State: state,
	}

	filename, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusOK, errors.ERROR_EXPORT_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errors.SUCCESS, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}

func ImportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, errors.ERROR, nil)
		return
	}

	tagService := tag_service.Tag{}
	err = tagService.Import(file)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusOK, errors.ERROR_IMPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errors.SUCCESS, nil)
}
