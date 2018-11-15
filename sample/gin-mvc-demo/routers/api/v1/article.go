package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"net/http"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/errors"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/setting"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/util"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/app"
	"github.com/Theodoree/gin/go-gin-example/service/article_service"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/service/tag_service"
	"github.com/Theodoree/gin/go-gin-example/pkg/e"
)

// @Summary 获取单个文章
// @Param id Param int  "ID"
// @Produce  json
// @Router /api/v1/articles/:id [Get]
func GetArticle(c *gin.Context) {
	AppG := app.Gin{c}
	id := com.StrTo(c.Param(`id`)).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, `id`).Message(`ID必须大于0`)
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil { //找不到文章
		AppG.Response(http.StatusOK, errors.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists { //
		AppG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	AppG.Response(http.StatusOK, errors.SUCCESS, article)
}

// @Summary 获取多个文章
// @Param state  Query int  "State"
// @Param tag_id Query int  "tagId"
// @Produce  json
// @Router /api/v1/articles [Get]
func GetArticles(c *gin.Context) {
	valid := validation.Validation{}
	AppG := app.Gin{c}
	state := com.StrTo(c.Query(`state`)).MustInt()
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		State:    state,
		TagID:    tagId,
		PageSize: setting.AppSetting.PageSize,
		PageNum:  util.Getpage(c),
	}

	ArticleList, err := articleService.GetAll()

	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	total, err := articleService.Count()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_COUNT_ARTICLE_FAIL, nil)
	}
	data := make(map[string]interface{})
	data["lists"] = ArticleList
	data["total"] = total
	AppG.Response(http.StatusOK, errors.SUCCESS, data)
}

// @Summary 新增文章
// @Param tag_id  Query int     "TagId"
// @Param title   Query string  "Title"
// @Param desc    Query string  "Desc"
// @Param content Query string  "Content"
// @Param created_by Query int  "CreatedBy"
// @Param state   Query int     "State"
// @Produce  json
// @Router /api/v1/articles [Get]
func AddArticle(c *gin.Context) {
	TagId := com.StrTo(c.Query("tag_id")).MustInt()
	Title := c.Query("title")
	Desc := c.Query("desc")
	Content := c.Query("content")
	CreatedBy := c.Query("created_by")
	State := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	AppG := app.Gin{c}
	valid := validation.Validation{}
	valid.Min(TagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(Title, "title").Message("标题不能为空")
	valid.Required(Desc, "desc").Message("简述不能为空")
	valid.Required(Content, "content").Message("内容不能为空")
	valid.Required(CreatedBy, "created_by").Message("创建人不能为空")
	valid.Range(State, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: TagId}
	exists, err := tagService.ExistByID()
	if err != nil {
		AppG.Response(http.StatusOK, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		AppG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	articleService := article_service.Article{
		TagID:     TagId,
		Title:     Title,
		Desc:      Desc,
		Content:   Content,
		CreatedBy: CreatedBy,
		State:     State,
	}

	if err := articleService.Add(); err != nil {
		AppG.Response(http.StatusOK, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	AppG.Response(http.StatusOK, errors.SUCCESS, nil)
}

// @Summary 修改文章
// @Param id  	  Param int     "Id"
// @Param tag_id  Query string  "TagId"
// @Param title   Query string  "Title"
// @Param desc 	  Query string  "Desc"
// @Param content Query int     "Content"
// @Param modified_by Query int "ModifiedBy"
// @Produce  json
// @Router /api/v1/articles/:id [Get]
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	Id := com.StrTo(c.Param("id")).MustInt()
	TagId := com.StrTo(c.Query("tag_id")).MustInt()
	Title := c.Query("title")
	Desc := c.Query("desc")
	Content := c.Query("content")
	ModifiedBy := c.Query("modified_by")
	AppG := app.Gin{c}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(Id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(Title, 100, "title").Message("标题最长为100字符")
	valid.Required(Title, "title").Message("标题不能为空")
	valid.MaxSize(Desc, 255, "desc").Message("简述最长为255字符")
	valid.Required(Desc, "desc").Message("简述不能为空")
	valid.MaxSize(Content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(Content, "content").Message("内容不能为空")
	valid.Required(ModifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(ModifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}

	tagService := tag_service.Tag{ID: TagId}
	Exist, err := tagService.ExistByID()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !Exist {
		AppG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	articleService := article_service.Article{
		ID:         Id,
		State:      state,
		TagID:      TagId,
		Title:      Title,
		Desc:       Desc,
		Content:    Content,
		ModifiedBy: ModifiedBy,
	}
	Exist, err = articleService.ExistByID()
	if err != nil {
		AppG.Response(http.StatusOK, errors.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !Exist {
		AppG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	AppG.Response(http.StatusOK, errors.SUCCESS, nil)
}

// @Summary 删除文章
// @Param id  	  Param int     "Id"
// @Produce  json
// @Router /api/v1/articles/:id [Get]
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	AppG := app.Gin{c}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		AppG.Response(http.StatusOK, errors.INVALID_PARAMS, nil)
		return
	}
	articleService:=article_service.Article{
		ID:id,
	}
	Exist,err:=articleService.ExistByID()
	if err != nil { //找不到文章
		AppG.Response(http.StatusOK, errors.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !Exist { //
		AppG.Response(http.StatusOK, errors.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err =articleService.Delete()
	if err !=nil{
		AppG.Response(http.StatusOK,errors.ERROR_DELETE_ARTICLE_FAIL,nil)
		return
	}
	AppG.Response(http.StatusOK, errors.SUCCESS, nil)
}
